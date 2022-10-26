package fixtures

import (
	"testing"
)

func TestRead(t *testing.T) {
	r, err := Read("../fixtures/argo/argoUbunto.yaml")
	if err != nil {
		t.Error(err)
	}
	println(string(r))
}
