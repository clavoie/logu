package logu

import (
	"log"
	"os"
)

// goLogger is an implementation of Logger that wraps the
// Go logging package:
//
// https://golang.org/pkg/log/
//
// Only calls to Print, Printf, and Println are made from this
// implementation. No calls to any of the Fatal or Panic functions
// are ever made.
//
// Info and Debug logs are written to Stdout, Warning and Error logs
// are written to Stderr. Log lines are written with a prefix of
// '[TYPE] ', such that Info logs have a prefix of '[INFO] ', error
// logs have a prefix of '[ERROR] ', etc. log.LstdFlags are used for
// all logs.
type goLogger struct {
	debugLogger *log.Logger
	errLogger   *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
}

// goLogger is an implementation of Logger that wraps the
// Go logging package:
//
// https://golang.org/pkg/log/
//
// Only calls to Print, Printf, and Println are made from this
// implementation. No calls to any of the Fatal or Panic functions
// are ever made.
//
// Info and Debug logs are written to Stdout, Warning and Error logs
// are written to Stderr. Log lines are written with a prefix of
// '[TYPE] ', such that Info logs have a prefix of '[INFO] ', error
// logs have a prefix of '[ERROR] ', etc. log.LstdFlags are used for
// all logs.
func NewGoLogger() Logger {
	return &goLogger{
		errLogger:   log.New(os.Stderr, "[ERROR] ", log.LstdFlags),
		debugLogger: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags),
		infoLogger:  log.New(os.Stdout, "[INFO] ", log.LstdFlags),
		warnLogger:  log.New(os.Stderr, "[WARN] ", log.LstdFlags),
	}
}

func (g *goLogger) Debugf(format string, args ...interface{}) {
	g.debugLogger.Printf(format, args...)
}

func (g *goLogger) Errorf(format string, args ...interface{}) {
	g.errLogger.Printf(format, args...)
}

func (g *goLogger) Infof(format string, args ...interface{}) {
	g.infoLogger.Printf(format, args...)
}

func (g *goLogger) Warningf(format string, args ...interface{}) {
	g.warnLogger.Printf(format, args...)
}
