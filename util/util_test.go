package util

import (
	"testing"
)

func TestGetAWS(t *testing.T) {

	r, err := GetAWS()
	if err != nil {
		t.Error(err)
	}
	_ = r
}
