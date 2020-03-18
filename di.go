package logu

import "github.com/clavoie/di/v2"

// NewAppEngineDiDefs returns di definitions suitable for a project
// run in Google's AppEngine.
func NewAppEngineDiDefs() []*di.Def {
	return []*di.Def{
		{NewAppEngineLogger, di.PerHttpRequest},
	}
}
