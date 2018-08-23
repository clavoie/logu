package logu

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// appEngineLogger represents a Logger implementation which logs to the Google App Engine
type appEngineLogger struct {
	context context.Context
}

// NewAppEngineLogger creates a new logging implementation from the current http request
// that logs to the Google AppEngine.
func NewAppEngineLogger(r *http.Request) Logger {
	return &appEngineLogger{
		context: appengine.NewContext(r),
	}
}

func (ael *appEngineLogger) Debugf(format string, args ...interface{}) {
	log.Debugf(ael.context, format, args...)
}

func (ael *appEngineLogger) Errorf(format string, args ...interface{}) {
	log.Errorf(ael.context, format, args...)
}

func (ael *appEngineLogger) Infof(format string, args ...interface{}) {
	log.Infof(ael.context, format, args...)
}

func (ael *appEngineLogger) Warningf(format string, args ...interface{}) {
	log.Warningf(ael.context, format, args...)
}
