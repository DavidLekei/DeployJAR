package main

import (
	"os"
	"strings"
)

var defaultConfigPath = "\\res\\config.cfg"

const MAX_LINE_LENGTH int = 1024
const NEWLINE_CHARACATER int = 10
const WHENCE_FILE_ORIGIN int = 0

type Config struct {
	defaultJarFilePath string
	defaultEnv         string
	verbose            bool
}

type LineReader struct {
	offset int64
	file   *os.File
}

type FileLineReader interface {
	ReadLine() (string, error)
}

func (reader *LineReader) ReadLine() (string, error) {

	reader.file.Seek(reader.offset, WHENCE_FILE_ORIGIN)

	stringBuffer := make([]byte, MAX_LINE_LENGTH)
	buffer := make([]byte, 1)

	n, err := reader.file.Read(buffer)
	var count int64

	for err == nil && buffer[0] != '\n' && n != 0 {
		stringBuffer[count] = buffer[0]
		count++
		n, err = reader.file.Read(buffer)
	}

	s := string(stringBuffer[0:count])

	reader.offset = reader.offset + count + 1 //+1 to go PAST the newline character

	return s, err
}

func LoadConfig() *map[string]string {
	return LoadConfigFromPath(defaultConfigPath)
}

func LoadConfigFromPath(configFilePath string) *map[string]string {
	dir, err := os.Getwd()
	file, err := os.Open(dir + configFilePath)
	if err != nil {
		panic("ERROR: Could not open file: " + dir + configFilePath)
	}

	reader := LineReader{
		offset: 0,
		file:   file,
	}

	config := make(map[string]string)

	line, err := reader.ReadLine()
	for err == nil {
		s := strings.Split(line, "=")
		config[s[0]] = s[1]
		line, err = reader.ReadLine()
	}

	return &config
}
