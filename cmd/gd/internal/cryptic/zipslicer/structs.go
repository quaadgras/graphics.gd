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

const (
	fileHeaderSignature      = 0x04034b50
	directoryHeaderSignature = 0x02014b50
	directoryEndSignature    = 0x06054b50
	directory64LocSignature  = 0x07064b50
	directory64EndSignature  = 0x06064b50
	dataDescriptorSignature  = 0x08074b50
	fileHeaderLen            = 30
	directoryHeaderLen       = 46
	directoryEndLen          = 22
	directory64LocLen        = 20
	directory64EndLen        = 56
	dataDescriptorLen        = 16
	dataDescriptor64Len      = 24
	zip64ExtraID             = 0x0001
	zip64ExtraLen            = 24

	zip20 = 20
	zip45 = 45

	uint32Max = 0xffffffff
	uint16Max = 0xffff
)

type zipCentralDir struct {
	Signature        uint32
	CreatorVersion   uint16
	ReaderVersion    uint16
	Flags            uint16
	Method           uint16
	ModifiedTime     uint16
	ModifiedDate     uint16
	CRC32            uint32
	CompressedSize   uint32
	UncompressedSize uint32
	FilenameLen      uint16
	ExtraLen         uint16
	CommentLen       uint16
	StartDisk        uint16
	InternalAttrs    uint16
	ExternalAttrs    uint32
	Offset           uint32
}

type zip64End struct {
	Signature      uint32
	RecordSize     uint64
	CreatorVersion uint16
	ReaderVersion  uint16
	Disk           uint32
	FirstDisk      uint32
	DiskCDCount    uint64
	TotalCDCount   uint64
	CDSize         uint64
	CDOffset       uint64
}

type zip64Loc struct {
	Signature uint32
	Disk      uint32
	Offset    uint64
	DiskCount uint32
}
type zipEndRecord struct {
	Signature     uint32
	DiskNumber    uint16
	DiskCD        uint16
	DiskCDCount   uint16
	TotalCDCount  uint16
	CDSize        uint32
	CDOffset      uint32
	CommentLength uint16
}

type zipLocalHeader struct {
	Signature        uint32
	ReaderVersion    uint16
	Flags            uint16
	Method           uint16
	ModifiedTime     uint16
	ModifiedDate     uint16
	CRC32            uint32
	CompressedSize   uint32
	UncompressedSize uint32
	FilenameLen      uint16
	ExtraLen         uint16
}

type zip64Extra struct {
	Signature        uint16
	RecordSize       uint16
	UncompressedSize uint64
	CompressedSize   uint64
	Offset           uint64
}

type zipDataDesc struct {
	Signature        uint32
	CRC32            uint32
	CompressedSize   uint32
	UncompressedSize uint32
}

type zipDataDesc64 struct {
	Signature        uint32
	CRC32            uint32
	CompressedSize   uint64
	UncompressedSize uint64
}
