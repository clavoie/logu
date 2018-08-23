package logu_test

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/clavoie/di"
	"github.com/clavoie/erru"
	"github.com/clavoie/logu"
)

func onResolveErr(err *di.ErrResolve, w http.ResponseWriter, r *http.Request) {
	logger := logu.NewAppEngineLogger(r)
	logger.Errorf("err encountered while resolving dependencies: %v", err.String())

	httpErr, isHttpErr := err.Err.(erru.HttpErr)
	if isHttpErr {
		w.WriteHeader(httpErr.StatusCode())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type Dependency interface {
	DoWork() bool
}

type dependency struct{}

func (d *dependency) DoWork() bool { return false }
func NewDependency() Dependency    { return new(dependency) }

var defs = []*di.Def{{NewDependency, di.PerHttpRequest}}

func MyHandler(logger logu.Logger, dep Dependency) {
	result := dep.DoWork()

	if result == false {
		logger.Errorf("Value for result: %v", result)
	}
}

func ExampleAppEngine() {
	resolver, err := di.NewResolver(onResolveErr, defs, logu.NewAppEngineDiDefs())

	if err != nil {
		log.Fatal(err)
	}

	httpHandler, err := resolver.HttpHandler(MyHandler)

	if err != nil {
		log.Fatal(err)
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	httpHandler(w, req)
}
