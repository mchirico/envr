package kustomize

import (
	"fmt"
	"github.com/mchirico/envr/fixtures"
	"testing"
)

func TestCreateKustomizeTmpDir(t *testing.T) {
	dir := fixtures.Path("../junkt")
	path, err := CreateKustomizeTmpDir(dir)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Println(path)
	err = PopulateKustomizeTmpDir(path)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
