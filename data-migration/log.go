package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type LogWriter struct {
	File   *os.File
	Logger *log.Logger
}

func NewLogWriter(dir string, prefix string) (*LogWriter, error) {
	f, err := createLogFile(dir)
	if err != nil {
		return nil, err
	}

	mw := io.MultiWriter(f, os.Stdout)

	l := log.New(mw, prefix, log.Ldate)

	return &LogWriter{
		File:   f,
		Logger: l,
	}, nil
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
