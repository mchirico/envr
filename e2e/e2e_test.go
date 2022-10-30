package e2e_test

import (
	"github.com/mchirico/envr/file"
	"github.com/mchirico/envr/fixtures"
	"testing"
)

func TestE2E(t *testing.T) {
	dir := fixtures.Path("./tmp/junk/")
	err := file.RemoveAll(dir)

	err = file.MkdirAll(dir)
	if err != nil {
		t.Errorf("Error making dir: %s", err)
	}
	file.Write(dir+"/file", []byte("test"))

	file.RemoveAll(dir)
}
