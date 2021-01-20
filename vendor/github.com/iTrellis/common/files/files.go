/*
Copyright © 2020 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package files

import (
	"io"
	"os"
	"sync"
)

// FileMode
const (
	FileModeOnlyRead  os.FileMode = 0444
	FileModeReadWrite os.FileMode = 0666
)

var defaultFile *fileGem

type fileGem struct {
	sync.Mutex

	executingPath map[string]FileStatus
	readBufLength int
}

type callbackExec func() error

// New return filerepo with default executor
func New() FileRepo {
	if defaultFile == nil {
		defaultFile = &fileGem{
			executingPath: make(map[string]FileStatus),
			readBufLength: ReadBufferLength,
		}
	}
	return defaultFile
}

func (p *fileGem) updateExecFileStatus(name string, status FileStatus) error {
	p.Lock()
	defer p.Unlock()
	if p.FileOpened(name) && status != FileStatusClosed {
		return ErrFileIsAlreadyOpen
	}
	if status == FileStatusClosed {
		delete(p.executingPath, name)
		return nil
	}
	p.executingPath[name] = status

	return nil
}

func (p *fileGem) Read(name string) (b []byte, n int, err error) {

	f := func() error {
		b, n, err = p.read(name, p.readBufLength)
		return err
	}

	err = p.execute(name, FileStatusOpening, f)

	return
}

func (p *fileGem) read(name string, bufLen int) (b []byte, n int, err error) {

	fi, e := p.tryOpen(name)
	if e != nil {
		err = e
		return
	}
	defer fi.Close()
	for {

		buf := make([]byte, bufLen)
		m, e := fi.Read(buf)
		if e != nil && e != io.EOF {
			err = ErrFailedReadFile
			return
		}
		n += m
		b = append(b, buf[:m]...)
		if m < bufLen {
			break
		}
	}

	return
}

func (p *fileGem) FileOpened(name string) bool {
	return p.executingPath[name] != FileStatusClosed
}

func (p *fileGem) tryOpen(name string) (*os.File, error) {
	return p.tryOpenfile(name, os.O_RDONLY, FileModeOnlyRead)
}

func (p *fileGem) tryOpenfile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (p *fileGem) Write(name, s string) (int, error) {
	return p.WriteBytes(name, []byte(s))
}

func (p *fileGem) WriteBytes(name string, b []byte) (int, error) {
	return p.write(name, b, os.O_TRUNC)
}

func (p *fileGem) WriteAppend(name, s string) (int, error) {
	return p.WriteAppendBytes(name, []byte(s))
}

func (p *fileGem) WriteAppendBytes(name string, b []byte) (int, error) {
	return p.write(name, b, os.O_APPEND)
}

func (p *fileGem) Rename(oldpath, newpath string) error {

	f := func() error {
		return os.Rename(oldpath, newpath)
	}

	return p.execute(oldpath, FileStatusMoving, f)
}

func (p *fileGem) SetReadBufLength(l int) error {
	if l <= 0 {
		return ErrReadBufferLengthBelowZero
	}

	p.readBufLength = l

	return nil
}

func (p *fileGem) write(name string, b []byte, flag int) (n int, err error) {

	callback := func() error {
		fi, e := p.tryOpenfile(name, os.O_CREATE|os.O_WRONLY|flag, FileModeReadWrite)
		if e != nil {
			return e
		}
		defer fi.Close()

		n, e = fi.Write(b)
		return e
	}

	err = p.execute(name, FileStatusOpening, callback)

	return
}

func (p *fileGem) execute(name string, fStatus FileStatus, callback callbackExec) (err error) {
	if err = p.updateExecFileStatus(name, fStatus); err != nil {
		return
	}
	defer func() { _ = p.updateExecFileStatus(name, FileStatusClosed) }()

	return callback()
}

func (p *fileGem) FileInfo(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
