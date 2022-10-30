package e2e_test

import (
	"github.com/mchirico/envr/file"
	"github.com/mchirico/envr/fixtures"
	"testing"
)

func SetupTmpArea(path string) (string, error) {
	dir := fixtures.Path(path)
	file.RemoveAll(dir)

	err := file.MkdirAll(dir)
	return dir, err
}

func RemoveTmp(path string) error {
	dir := fixtures.Path(path)
	return file.RemoveAll(dir)
}

func TestE2E(t *testing.T) {
	path := "./tmp/junk/"
	if dir, err := SetupTmpArea(path); err != nil {
		t.Error(err)
		return
	} else {
		file.Write(dir+"/file", []byte("test"))

	}

	RemoveTmp(path)
}
