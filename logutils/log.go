package logutils

import (
	"io"

	"log"
)

type EmptyWriter struct{}

func (e EmptyWriter) Write(p []byte) (n int, err error) { return }

func NullLogOutput() {
	log.SetOutput(EmptyWriter{})
}

type LogWriter struct {
	Logger *log.Logger
}

func (l *LogWriter) Write(p []byte) (n int, err error) {
	l.Logger.Println(string(p))
	return len(p), nil
}

func ToWriter(l *log.Logger) io.Writer {
	return &LogWriter{l}
}
