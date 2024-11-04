package logger

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

const numberOfStackFrameSkip = 2

// DefaultLogger return configured default logger
func DefaultLogger() *logrus.Logger {
	return logger
}

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields logrus.Fields

// SetLogLevel ...
func SetLogLevel(level logrus.Level) {
	logger.Level = level
}

// SetLogFormatter ...
func SetLogFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
		entry.Debug(args...)
	}
}

// DebugWithFields Debug logs a message with fields at level Debug on the standard logger
func DebugWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
		entry.Debug(l)
	}
}

// Println Info logs a message at level Info on the standard logger
func Println(args ...interface{}) {
	Info(args...)
}

// Info logs a message at level Info on the standard logger
func Info(args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
		entry.Info(args...)
	}
}

// InfoWithFields logs a message with fields at level Info on the standard logger
func InfoWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.InfoLevel {
		entry := logrus.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
		entry.Info(l)
	}
}

// Warn logs a message at level Warn on the standard logger
func Warn(args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logrus.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
		entry.Warn(args...)
	}
}

// WarnWithFields logs a message with fields at level Warn on the standard logger
func WarnWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
		entry.Warn(l)
	}
}

// Error logs a message at level Error on the standard logger
func Error(args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logrus.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
		entry.Error(args...)
	}
}

// ErrorWithFields logs a message with fields at level Error on the standard logger
func ErrorWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
		entry.Error(l)
	}
}

// Fatal logs a message at level Fatal on the standard logger
func Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
		entry.Fatal(args...)
	}
}

// FatalWithFields logs a message with fields at level Fatal on the standard logger
func FatalWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
		entry.Fatal(l)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
	entry.Panic(args...)
}

// PanicWithFields logs a message with fields at level Panic on the standard logger
func PanicWithFields(l interface{}, f Fields) {
	entry := logger.WithFields(logrus.Fields(f))
	entry.Data["file"] = fileInfo(numberOfStackFrameSkip)
	entry.Panic(l)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}

	return fmt.Sprintf("%s:%d", file, line)
}
