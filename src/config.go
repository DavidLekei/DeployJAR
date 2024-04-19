package main

import (
	"fmt"
	"os"
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

func (reader LineReader) ReadLine() (string, error) {

	reader.file.Seek(reader.offset, WHENCE_FILE_ORIGIN)

	stringBuffer := make([]byte, MAX_LINE_LENGTH)
	buffer := make([]byte, 1)

	n, err := reader.file.Read(buffer)
	count := 0

	for err == nil {
		stringBuffer[count] = buffer[0]
		count++
	}

}

func LoadConfig() *Config {
	return LoadConfigFromPath(defaultConfigPath)
}

func LoadConfigFromPath(configFilePath string) *Config {
	dir, err := os.Getwd()
	file, err := os.Open(dir + configFilePath)
	if err != nil {
		panic("ERROR: Could not open file: " + dir + configFilePath)
	}

	fmt.Println("File contents: ", file)

	reader := LineReader{
		offset: 0,
		file:   file,
	}

	line, err := reader.ReadLine()

	for err == nil {
		line, err = reader.ReadLine()
		fmt.Println("Line: ", line)
	}

	return nil
}
