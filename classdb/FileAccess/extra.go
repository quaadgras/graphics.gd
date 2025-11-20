package FileAccess

import (
	"io"
	"io/fs"
	"runtime"
	"time"

	"graphics.gd/variant/Error"
)

func tryConvertToFS(err error) error {
	if err == nil {
		return nil
	}
	switch err := err.(type) {
	case Error.Code:
		switch err {
		case Error.FileNotFound:
			return fs.ErrNotExist
		case Error.FileEof:
			return io.EOF
		case Error.FileNoPermission:
			return fs.ErrPermission
		case Error.FileAlreadyInUse:
			return fs.ErrExist
		default:
			return err
		}
	default:
		return err
	}
}

type fileInfo struct {
	name    string
	size    int64
	mode    fs.FileMode
	modTime time.Time
	sys     Instance
}

func (fi fileInfo) Name() string           { return fi.name }
func (fi fileInfo) Size() int64            { return fi.size }
func (fi fileInfo) Mode() fs.FileMode      { return fi.mode }
func (fi fileInfo) ModTime() (t time.Time) { return fi.modTime }
func (fi fileInfo) IsDir() bool            { return false }
func (fi fileInfo) Sys() any               { return fi.sys }

func (self Instance) Stat() (fs.FileInfo, error) {
	var info fileInfo
	info.name = self.GetPath()
	if err := self.GetError(); err != nil {
		return nil, tryConvertToFS(err)
	}
	info.size = int64(self.GetLength())
	if err := self.GetError(); err != nil {
		return nil, tryConvertToFS(err)
	}
	switch runtime.GOOS {
	case "windows":
		info.mode = 0o666
	default:
		info.mode = fs.FileMode(GetUnixPermissions(self.GetPath()))
		if err := self.GetError(); err != nil {
			return nil, tryConvertToFS(err)
		}
	}
	info.modTime = time.Unix(int64(GetModifiedTime(self.GetPath())), 0)
	if err := self.GetError(); err != nil {
		return nil, tryConvertToFS(err)
	}
	info.sys = self
	return info, nil
}

func (self Instance) Read(dst []byte) (int, error) {
	src := self.GetBuffer(len(dst))
	n := copy(dst, src)
	if err := self.GetError(); err != nil {
		return n, tryConvertToFS(err)
	}
	return n, nil
}

func (self Instance) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		self.SeekTo(int(offset))
	case io.SeekCurrent:
		self.SeekTo(int(self.GetPosition() + int(offset)))
	case io.SeekEnd:
		self.SeekTo(int(self.GetLength() + int(offset)))
	}
	return int64(self.GetPosition()), tryConvertToFS(self.GetError())
}

func (self Instance) Write(src []byte) (int, error) {
	if ok := self.StoreBuffer(src); !ok {
		return 0, tryConvertToFS(self.GetError())
	}
	return len(src), nil
}

func (self Instance) Close() error {
	Advanced(self).Close()
	return tryConvertToFS(self.GetError())
}
