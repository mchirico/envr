package e2e_test

import (
	"github.com/mchirico/envr/configuration"
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

func SampleFile() string {
	return `# Sample File
key={{TOKEN}}
key2=value2
`
}

func TestE2E(t *testing.T) {
	path := "./tmp/junk/"
	if dir, err := SetupTmpArea(path); err != nil {
		t.Error(err)
		return
	} else {
		configuration.SetPath(fixtures.Path(".envr"))
		e, err := configuration.ProfileEnvExports("argo")
		if err != nil {
			t.Errorf("Error getting exports: %s", err)
		}
		if v, ok := e["key"]; !ok {
			t.Errorf("Expected key, got %s", v)
			return
		} else {
			t.Logf("key: %s", v)
			r := file.StringReplaceAll(SampleFile(), e["replace_string"], e["value"])
			if err := file.Write(dir+"/test.txt", []byte(r)); err != nil {
				t.Error(err)
				return
			}
		}

		r, _ := file.Read(dir + "/test.txt")
		t.Logf("r: %s", r)

	}

	RemoveTmp(path)
}
