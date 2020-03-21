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
//
// Log lines are written with a prefix of
// '[TYPE] ', such that Info logs have a prefix of '[INFO] ', error
// logs have a prefix of '[ERROR] ', etc. log.LstdFlags are used for
// all logs.
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
//
// Log lines are written with a prefix of
// '[TYPE] ', such that Info logs have a prefix of '[INFO] ', error
// logs have a prefix of '[ERROR] ', etc. log.LstdFlags are used for
// all logs.
func NewLogPkgLogger() Logger {
	return &logPkgLogger{}
}

func (l *logPkgLogger) Debugf(format string, args ...interface{}) {
	log.Printf("[DEBUG] "+format, args...)
}

func (l *logPkgLogger) Errorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

func (l *logPkgLogger) Infof(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func (l *logPkgLogger) Warningf(format string, args ...interface{}) {
	log.Printf("[WARN] "+format, args...)
}
