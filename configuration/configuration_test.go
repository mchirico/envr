package configuration

import (
	"fmt"
	"github.com/mchirico/envr/fixtures"
	"testing"
)

func TestProfileEnvExports(t *testing.T) {
	err := SetPath(fixtures.Path(".envr"))
	if err != nil {
		t.Errorf("Error setting path: %s", err)
	}
	e, err := ProfileEnvExports("argo")
	if err != nil {
		t.Errorf("Error getting exports: %s", err)
	}
	if v, ok := e["key"]; !ok {
		t.Errorf("Expected key, got %s", v)
	}
	fmt.Println(e)

}

func Test_ListAllProfiles(t *testing.T) {
	err := SetPath(fixtures.Path(".envr"))
	if err != nil {
		t.Errorf("Error setting path: %s", err)
	}
	profiles := ProfilesAvailable()
	count := 0
	for _, v := range profiles {
		if v == "argo" {
			count++
		}
		t.Logf("%s", v)
	}
	if count != 1 {
		t.Errorf("Expected 1 profile, got %d", count)
	}
}
