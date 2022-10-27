package pb

import (
	"testing"
)

func Test_writeAll(t *testing.T) {
	Write("cpy test..abc")
	s, err := Read()
	if err != nil {
		t.Error(err)
	}
	if s != "cpy test..abc" {
		t.Errorf("expected cpy test..abc, got %s", s)
	}

}
