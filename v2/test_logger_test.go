package logu_test

import (
	"testing"

	"github.com/clavoie/logu/v2"
)

func TestTestLogger(t *testing.T) {
	logger := logu.NewTestLogger(t)
	logger.Debugf("debug: %v", 1)
	// logger.Errorf("error: %v", 2)
	logger.Infof("info: %v", 3)
	logger.Warningf("warning: %v", 4)
}
