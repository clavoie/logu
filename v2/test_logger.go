package logu

import "testing"

// TestLogger is an ILogger implementation that logs for tests
type testLogger struct {
	t *testing.T
}

// NewTestLogger returns a new Logger implementation that logs for tests.
func NewTestLogger(t *testing.T) Logger {
	return &testLogger{t}
}

func (tl *testLogger) Debugf(format string, args ...interface{}) {
	tl.t.Logf(format, args...)
}

func (tl *testLogger) Errorf(format string, args ...interface{}) {
	tl.t.Errorf(format, args...)
}

func (tl *testLogger) Infof(format string, args ...interface{}) {
	tl.t.Logf(format, args...)
}

func (tl *testLogger) Warningf(format string, args ...interface{}) {
	tl.t.Logf(format, args...)
}
