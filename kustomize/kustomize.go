package kustomize

import (
	"fmt"
	"github.com/mchirico/envr/fixtures"
	"os"
	"path/filepath"
	"strings"
)

func CreateKustomizeTmpDir(dir string) (string, error) {
	newpath := filepath.Join(dir, "base")
	err := os.MkdirAll(newpath, os.ModePerm)
	return newpath, err
}

func PopulateKustomizeTmpDir(dir string) error {
	files := []string{
		fixtures.Path("./kustomize/base/kustomization.yaml"),
		fixtures.Path("./kustomize/base/pod.yaml"),
		fixtures.Path("./kustomize/base/patch.yaml")}

	for _, f := range files {
		err := populateKustomizeTmpDir(dir, f)
		if err != nil {
			return err
		}
	}

	return nil

}
func populateKustomizeTmpDir(dir string, filefrom string) error {
	file, err := fixtures.Read(filefrom)
	if err != nil {
		return err
	}
	f := strings.SplitN(filefrom, "/", -1)
	l := len(f)
	if l == 0 {
		return fmt.Errorf("filefrom: %s", filefrom)
	}
	err = fixtures.Write(filepath.Join(dir, f[l-1]), file)
	if err != nil {
		return err
	}
	return nil
}
