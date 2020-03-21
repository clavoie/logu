package logu

import (
	"log"
)

// logPkgLogger is an implementation of Logger that uses the
// Go logging package to write to the standard logger:
//
// https://golang.org/pkg/log/
//
// Only calls to Print, Printf, and Println are made from this
// implementation. No calls to any of the Fatal or Panic functions
// are ever made.
type logPkgLogger struct {
}

// NewLogPkgLogger creates an an implementation of Logger that uses the
// Go logging package to write to the standard logger:
//
// https://golang.org/pkg/log/
//
// Only calls to Print, Printf, and Println are made from this
// implementation. No calls to any of the Fatal or Panic functions
// are ever made.
func NewLogPkgLogger() Logger {
	return &logPkgLogger{}
}

func (l *logPkgLogger) Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func (l *logPkgLogger) Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func (l *logPkgLogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *logPkgLogger) Warningf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
