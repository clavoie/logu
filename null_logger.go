package logu

// nullLogger represents a logging implementation that does nothing
type nullLogger struct {
}

// NewNullLogger returns a new Logger implementation that does nothing.
// All logging messages are destroyed.
func NewNullLogger() Logger {
	return &nullLogger{}
}

func (nl *nullLogger) Debugf(format string, args ...interface{}) {
}

func (nl *nullLogger) Errorf(format string, args ...interface{}) {
}

func (nl *nullLogger) Infof(format string, args ...interface{}) {
}

func (nl *nullLogger) Warningf(format string, args ...interface{}) {
}
