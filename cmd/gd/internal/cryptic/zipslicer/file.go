//
// Copyright (c) SAS Institute Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package zipslicer

import (
	"archive/zip"
	"bufio"
	"bytes"
	"compress/flate"
	"crypto"
	"encoding/binary"
	"errors"
	"hash"
	"hash/crc32"
	"io"
	"io/ioutil"
	"time"
)

type File struct {
	CreatorVersion   uint16
	ReaderVersion    uint16
	Flags            uint16
	Method           uint16
	ModifiedTime     uint16
	ModifiedDate     uint16
	CRC32            uint32
	CompressedSize   uint64
	UncompressedSize uint64
	Name             string
	Extra            []byte
	Comment          []byte
	InternalAttrs    uint16
	ExternalAttrs    uint32
	Offset           uint64

	r                 io.ReaderAt
	rs                int64
	raw               []byte
	lfh               zipLocalHeader
	lfhName, lfhExtra []byte
	ddb               []byte
	compd             []byte
}

type Reader struct {
	f     *File
	rc    io.ReadCloser
	crc   hash.Hash32
	err   error
	nread uint64
}

func (f *File) readLocalHeader() error {
	if f.lfh.Signature != 0 {
		// already done
		return nil
	}
	sr := io.NewSectionReader(f.r, int64(f.Offset), f.rs)
	var lfhb [fileHeaderLen]byte
	if _, err := io.ReadFull(sr, lfhb[:]); err != nil {
		return err
	}
	_ = binary.Read(bytes.NewReader(lfhb[:]), binary.LittleEndian, &f.lfh)
	if f.lfh.Signature != fileHeaderSignature {
		return errors.New("local file header not found")
	}
	f.lfhName = make([]byte, f.lfh.FilenameLen)
	if _, err := io.ReadFull(sr, f.lfhName); err != nil {
		return err
	}
	f.lfhExtra = make([]byte, f.lfh.ExtraLen)
	if _, err := io.ReadFull(sr, f.lfhExtra); err != nil {
		return err
	}
	return nil
}

func (f *File) readDataDesc() error {
	if err := f.readLocalHeader(); err != nil {
		return err
	}
	if f.lfh.Flags&0x8 == 0 {
		// no descriptor
		return nil
	}
	if len(f.ddb) != 0 {
		// already done
		return nil
	}
	lfhSize := fileHeaderLen + len(f.lfhName) + len(f.lfhExtra)
	pos := int64(f.Offset) + int64(lfhSize) + int64(f.CompressedSize)
	f.ddb = make([]byte, dataDescriptor64Len)
	if _, err := f.r.ReadAt(f.ddb[:dataDescriptorLen], pos); err != nil {
		return err
	}
	// Read the 32-bit len so we don't overshoot. The underlying stream might
	// not be seekable.
	var desc zipDataDesc
	_ = binary.Read(bytes.NewReader(f.ddb[:dataDescriptorLen]), binary.LittleEndian, &desc)
	if desc.Signature != dataDescriptorSignature {
		return errors.New("data descriptor signature is missing")
	}
	if f.UncompressedSize >= uint32Max || desc.UncompressedSize != uint32(f.UncompressedSize) || desc.CompressedSize != uint32(f.CompressedSize) {
		// 64-bit
		if _, err := f.r.ReadAt(f.ddb[dataDescriptorLen:], pos+dataDescriptorLen); err != nil {
			return err
		}
		var desc64 zipDataDesc64
		_ = binary.Read(bytes.NewReader(f.ddb), binary.LittleEndian, &desc64)
		if desc64.CompressedSize != f.CompressedSize || desc64.UncompressedSize != f.UncompressedSize {
			return errors.New("data descriptor is invalid")
		}
		f.CRC32 = desc64.CRC32
	} else {
		// 32-bit
		f.ddb = f.ddb[:dataDescriptorLen]
		f.CRC32 = desc.CRC32
	}
	return nil
}

func (f *File) GetDirectoryHeader() ([]byte, error) {
	if len(f.raw) > 0 {
		return f.raw, nil
	}
	hdr := zipCentralDir{
		Signature:        directoryHeaderSignature,
		CreatorVersion:   f.CreatorVersion,
		ReaderVersion:    f.ReaderVersion,
		Flags:            f.Flags,
		Method:           f.Method,
		ModifiedTime:     f.ModifiedTime,
		ModifiedDate:     f.ModifiedDate,
		CRC32:            f.CRC32,
		CompressedSize:   uint32(f.CompressedSize),
		UncompressedSize: uint32(f.UncompressedSize),
		InternalAttrs:    f.InternalAttrs,
		ExternalAttrs:    f.ExternalAttrs,
		Offset:           uint32(f.Offset),
		FilenameLen:      uint16(len(f.Name)),
		ExtraLen:         uint16(len(f.Extra)),
		CommentLen:       uint16(len(f.Comment)),
	}
	if f.CompressedSize >= uint32Max || f.UncompressedSize >= uint32Max || f.Offset >= uint32Max {
		hdr.CompressedSize = uint32Max
		hdr.UncompressedSize = uint32Max
		hdr.Offset = uint32Max
		extra := zip64Extra{
			Signature:        zip64ExtraID,
			RecordSize:       zip64ExtraLen,
			UncompressedSize: f.UncompressedSize,
			CompressedSize:   f.CompressedSize,
			Offset:           f.Offset,
		}
		b := bytes.NewBuffer(make([]byte, 0, zip64ExtraLen+4+len(f.Extra)))
		_ = binary.Write(b, binary.LittleEndian, extra)
		b.Write(f.Extra)
		f.Extra = b.Bytes()
		hdr.ExtraLen = uint16(b.Len())
		hdr.ReaderVersion = zip45
	}
	b := bytes.NewBuffer(make([]byte, 0, directoryHeaderLen+len(f.Name)+len(f.Extra)+len(f.Comment)))
	_ = binary.Write(b, binary.LittleEndian, hdr)
	b.WriteString(f.Name)
	b.Write(f.Extra)
	b.Write(f.Comment)
	return b.Bytes(), nil
}

func (f *File) GetLocalHeader() ([]byte, error) {
	if err := f.readLocalHeader(); err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(make([]byte, 0, fileHeaderLen+len(f.lfhName)+len(f.lfhExtra)))
	_ = binary.Write(b, binary.LittleEndian, f.lfh)
	b.Write(f.lfhName)
	b.Write(f.lfhExtra)
	return b.Bytes(), nil
}

func (f *File) GetDataDescriptor() ([]byte, error) {
	if err := f.readDataDesc(); err != nil {
		return nil, err
	}
	return f.ddb, nil
}

func (f *File) GetTotalSize() (int64, error) {
	if err := f.readDataDesc(); err != nil {
		return 0, err
	}
	return fileHeaderLen + int64(len(f.lfhName)+len(f.lfhExtra)+len(f.ddb)) + int64(f.CompressedSize), nil
}

func (d *Directory) NewFile(name string, extra, contents []byte, w io.Writer, mtime time.Time, deflate, useDesc bool) (*File, error) {
	var zh zip.FileHeader
	// need the side effect of this conversion
	zh.SetModTime(mtime) //nolint:staticcheck
	var fb bytes.Buffer
	method := zip.Deflate
	if deflate {
		c, err := flate.NewWriter(&fb, 9)
		if err != nil {
			return nil, err
		}
		if _, err := c.Write(contents); err != nil {
			return nil, err
		}
		if err := c.Close(); err != nil {
			return nil, err
		}
	} else {
		method = zip.Store
		fb.Write(contents)
	}
	crc := crc32.NewIEEE()
	crc.Write(contents)
	sum := crc.Sum32()
	buf := bufio.NewWriter(w)

	f := &File{
		CreatorVersion:   zip45,
		ReaderVersion:    zip20,
		Method:           method,
		ModifiedTime:     zh.ModifiedTime,
		ModifiedDate:     zh.ModifiedDate,
		CRC32:            sum,
		CompressedSize:   uint64(fb.Len()),
		UncompressedSize: uint64(len(contents)),
		Name:             name,
		Extra:            extra,

		lfhName:  []byte(name),
		lfhExtra: extra,
		compd:    fb.Bytes(),
	}
	if useDesc {
		f.Flags = 0x8
		f.ReaderVersion = zip45
	}
	f.lfh = zipLocalHeader{
		Signature:     fileHeaderSignature,
		ReaderVersion: f.ReaderVersion,
		Flags:         f.Flags,
		Method:        f.Method,
		ModifiedTime:  f.ModifiedTime,
		ModifiedDate:  f.ModifiedDate,
		FilenameLen:   uint16(len(name)),
		ExtraLen:      uint16(len(extra)),
	}
	if !useDesc {
		f.lfh.CRC32 = f.CRC32
		f.lfh.CompressedSize = uint32(f.CompressedSize)
		f.lfh.UncompressedSize = uint32(f.UncompressedSize)
	}
	if err := binary.Write(buf, binary.LittleEndian, f.lfh); err != nil {
		return nil, err
	}
	if _, err := buf.WriteString(name); err != nil {
		return nil, err
	}
	if _, err := buf.Write(extra); err != nil {
		return nil, err
	}
	if _, err := buf.Write(fb.Bytes()); err != nil {
		return nil, err
	}
	if useDesc {
		desc := zipDataDesc64{
			Signature:        dataDescriptorSignature,
			CRC32:            sum,
			CompressedSize:   uint64(fb.Len()),
			UncompressedSize: uint64(len(contents)),
		}
		ddb := bytes.NewBuffer(make([]byte, 0, dataDescriptor64Len))
		_ = binary.Write(ddb, binary.LittleEndian, desc)
		f.ddb = ddb.Bytes()
		if _, err := buf.Write(f.ddb); err != nil {
			return nil, err
		}
	}
	if err := buf.Flush(); err != nil {
		return nil, err
	}
	return d.AddFile(f)
}

func (f *File) ModTime() time.Time {
	fh := zip.FileHeader{ModifiedDate: f.ModifiedDate, ModifiedTime: f.ModifiedTime}
	// need the side effect of this conversion
	return fh.ModTime() //nolint:staticcheck
}

func (f *File) Open() (io.ReadCloser, error) {
	return f.OpenAndTeeRaw(nil)
}

// Open zip file for reading, but also write raw zip data to 'sink'.
func (f *File) OpenAndTeeRaw(sink io.Writer) (*Reader, error) {
	if err := f.readLocalHeader(); err != nil {
		return nil, err
	}
	var r io.Reader
	if f.compd != nil {
		r = bytes.NewReader(f.compd)
	} else {
		pos := int64(f.Offset) + fileHeaderLen + int64(f.lfh.FilenameLen) + int64(f.lfh.ExtraLen)
		r = io.NewSectionReader(f.r, pos, int64(f.CompressedSize))
	}
	if sink != nil {
		r = io.TeeReader(r, sink)
	}
	crc := crc32.NewIEEE()
	var rc io.ReadCloser
	switch f.Method {
	case zip.Store:
		rc = ioutil.NopCloser(r)
	case zip.Deflate:
		rc = flate.NewReader(r)
	default:
		return nil, errors.New("unsupported zip compression")
	}
	return &Reader{f: f, rc: rc, crc: crc}, nil
}

func (r *Reader) Read(d []byte) (int, error) {
	if r.err != nil {
		return 0, r.err
	}
	n, err := r.rc.Read(d)
	r.crc.Write(d[:n])
	r.nread += uint64(n)
	if err == nil {
		return n, nil
	}
	if err != io.EOF {
		r.err = err
		return n, err
	}
	if r.nread != r.f.UncompressedSize {
		return 0, io.ErrUnexpectedEOF
	}
	if r.f.lfh.Flags&0x8 != 0 {
		if err2 := r.f.readDataDesc(); err2 != nil {
			if err2 == io.EOF {
				err = io.ErrUnexpectedEOF
			} else {
				err = err2
			}
		}
	}
	if r.f.CRC32 != 0 && r.crc.Sum32() != r.f.CRC32 {
		err = zip.ErrChecksum
	}
	r.err = err
	return n, err
}

func (r *Reader) Close() error {
	return r.rc.Close()
}

func (f *File) Digest(hash crypto.Hash) ([]byte, error) {
	fc, err := f.Open()
	if err != nil {
		return nil, err
	}
	d := hash.New()
	if _, err := io.Copy(d, fc); err != nil {
		return nil, err
	}
	return d.Sum(nil), fc.Close()
}

// Dump the local header and contents of this file to a writer
func (f *File) Dump(w io.Writer) (int64, error) {
	lfh, err := f.GetLocalHeader()
	if err != nil {
		return 0, err
	}
	if _, err := w.Write(lfh); err != nil {
		return 0, err
	}
	if f.compd != nil {
		if _, err := w.Write(f.compd); err != nil {
			return 0, err
		}
	} else {
		pos := int64(f.Offset) + fileHeaderLen + int64(f.lfh.FilenameLen) + int64(f.lfh.ExtraLen)
		r := io.NewSectionReader(f.r, pos, int64(f.CompressedSize))
		if _, err := io.Copy(w, r); err != nil {
			return 0, err
		}
	}
	ddb, err := f.GetDataDescriptor()
	if err != nil {
		return 0, err
	}
	if _, err := w.Write(ddb); err != nil {
		return 0, err
	}
	return f.GetTotalSize()
}
