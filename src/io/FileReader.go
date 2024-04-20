package io

import (
	"os"
)

const MAX_LINE_LENGTH int = 1024
const NEWLINE_CHARACATER int = 10
const WHENCE_FILE_ORIGIN int = 0

type LineReader struct {
	Offset int64
	File   *os.File
}

type FileLineReader interface {
	ReadLine() (string, error)
}

func (reader *LineReader) ReadLine() (string, error) {

	reader.File.Seek(reader.Offset, WHENCE_FILE_ORIGIN)

	stringBuffer := make([]byte, MAX_LINE_LENGTH)
	buffer := make([]byte, 1)

	n, err := reader.File.Read(buffer)
	var count int64

	for err == nil && buffer[0] != '\n' && n != 0 {
		stringBuffer[count] = buffer[0]
		count++
		n, err = reader.File.Read(buffer)
	}

	s := string(stringBuffer[0:count])

	reader.Offset = reader.Offset + count + 1 //+1 to go PAST the newline character

	return s, err
}
