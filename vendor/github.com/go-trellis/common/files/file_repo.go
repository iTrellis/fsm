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

import "os"

// FileRepo execute file functions
type FileRepo interface {
	// judge if file is opening
	FileOpened(string) bool
	// read file
	Read(string) (b []byte, n int, err error)
	// rewrite file with context
	Write(name, context string) (int, error)
	WriteBytes(name string, b []byte) (int, error)
	// append context to the file
	WriteAppend(name, context string) (int, error)
	WriteAppendBytes(name string, b []byte) (int, error)
	// rename file
	Rename(oldpath, newpath string) error
	// set length of buffer to read file, default: 1024
	SetReadBufLength(int) error
	// get information with file name
	FileInfo(name string) (os.FileInfo, error)
}

// FileStatus defile file status
type FileStatus int

// file status
const (
	// nothing
	FileStatusClosed FileStatus = iota
	// file is opened
	FileStatusOpening
	// file is moving or rename
	FileStatusMoving
)

// ReadBufferLength default reader buffer length
const (
	ReadBufferLength = 1024
)
