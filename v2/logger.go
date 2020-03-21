package logu

// Logger represents a general interface around the logging implementation
type Logger interface {
	// Debugf prints a formatted message intended for debugging.
	Debugf(format string, args ...interface{})

	// Errorf prints a formatted error message to the log.
	Errorf(format string, args ...interface{})

	// Infof prints a formatted info message to the log.
	Infof(format string, args ...interface{})

	// Warningf prints a formatted warning message to the log.
	Warningf(format string, args ...interface{})
}

//go:generate mockgen -destination mock_logu\mocks.go github.com/clavoie/logu/v2 Logger
