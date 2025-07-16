package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type LogWriter struct {
	file   *os.File
	writer *bufio.Writer
}

func NewLogger(dir string, prefix string) (*log.Logger, error) {
	f, err := createLogFile(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	mw := io.MultiWriter(os.Stdout, f)

	return log.New(mw, prefix, log.Ldate), nil
}

func createLogFile(dir string) (*os.File, error) {
	pwd, _ := os.Getwd()
	path := pwd + dir
	if err := os.MkdirAll(path, 0700); err != nil {
		return nil, err
	}

	ts := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("log_%s.log", ts)
	filePath := filepath.Join(path, fileName)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return nil, err
	}
	log.Printf("Create log file at %s", filePath)

	return file, nil
}
