package yal

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Logger struct {
	Level     LogLevel
	Writer    io.Writer
	ErrWriter io.Writer
	mu        sync.Mutex
	Colors    bool
}

func New(lvl LogLevel, writer io.Writer, errWriter io.Writer, useColors bool) *Logger {
	return &Logger{
		Level:     lvl,
		Writer:    writer,
		ErrWriter: errWriter,
		mu:        sync.Mutex{},
		Colors:    useColors,
	}
}

func Default() *Logger {
	return New(LevelInfo, os.Stdout, os.Stderr, true)
}

func (l *Logger) Info(v ...interface{}) {
	l.write(LevelInfo, v)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.writef(LevelInfo, format, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.write(LevelDebug, v)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.writef(LevelDebug, format, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.write(LevelWarn, v)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.writef(LevelWarn, format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.write(LevelError, v)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.writef(LevelError, format, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.write(LevelFatal, v)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.writef(LevelFatal, format, v...)
	os.Exit(1)
}

func (l *Logger) write(level LogLevel, v []interface{}) {
	if level < l.Level {
		return
	}
	l.writef(level, "%s", v...)
}

func (l *Logger) writef(level LogLevel, format string, v ...interface{}) {
	if level < l.Level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	var timestamp = fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))

	var line = fmt.Sprintf("%s%s: [%s] %s%s\n", level.ColorStart(), timestamp, level.String(), fmt.Sprintf(format, v...), level.ColorEnd())
	if !l.Colors {
		line = fmt.Sprintf("%s: [%s] %s\n", timestamp, level.String(), fmt.Sprintf(format, v...))
	}

	// Write the log entry followed by a newline.
	if l.Level == LevelError || l.Level == LevelFatal {
		l.ErrWriter.Write([]byte(line))
	} else {
		l.Writer.Write([]byte(line))
	}
}
