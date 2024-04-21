package io

import (
	"io"
	"os"
)

func CopyFile(srcPath string, destPath string) {
	file, err := os.Open(srcPath)
	if err != nil {
		panic("FileCopier.go - CopyFile() - Unable to open file: " + srcPath)
	}

	destFile, err := os.Create(destPath)
	if err != nil {
		panic("FileCopier.go - CopyFile() - Unable to create file: " + destPath)
	}

	_, err = io.Copy(destFile, file)
	if err != nil {
		panic("FileCopier.go - CopyFile - Error during io.Copy()")
	}
}
