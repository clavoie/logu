package logu_test

import (
	"testing"

	"github.com/clavoie/logu"
)

func TestNewDiDefs(t *testing.T) {
	defs := logu.NewAppEngineDiDefs()

	if defs == nil {
		t.Fatal("Expecting non-nil defs")
	}

	defs2 := logu.NewAppEngineDiDefs()
	if defs[0] == defs2[0] {
		t.Fatal("Not expecting defs to match")
	}
}
