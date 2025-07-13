package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type LogWriter struct {
	file   *os.File
	writer *bufio.Writer
}

func NewLogWrite(dir string) (*LogWriter, error) {
	file, err := createLogFile(dir)
	if err != nil {
		return nil, err
	}

	return &LogWriter{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func (w *LogWriter) Write(r *result) error {
	ts := time.Now().Format("2006-01-02 15:04:05.000")

	status := "SUCCESS"
	if r.err != nil {
		status = "FAILED"
	}

	entry := fmt.Sprintf("[%s] %s: %s\n", ts, status, r.ToString())

	if _, err := w.writer.WriteString(entry); err != nil {
		return err
	}

	return w.writer.Flush() // write buffered data into file
}

func (w *LogWriter) Close() error {
	if err := w.writer.Flush(); err != nil {
		return err
	}
	return w.file.Close()
}
func createLogFile(dir string) (*os.File, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	ts := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("log_%s.log", ts)
	filePath := filepath.Join(dir, fileName)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return nil, err
	}
	log.Printf("Create log file at %s", filePath)

	return file, nil
}
