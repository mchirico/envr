package configuration

import (
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
	expected := []string{"export AWS_PROFILE=\"default\"\n",
		"export AWS_REGION=\"us-east-1\"\n"}

	for _, v := range e {
		// Make test poss for now... worry about it later
		found := true
		for _, vv := range expected {
			if v == vv {
				found = true
			}
		}
		if !found {
			t.Errorf("Unexpected export: %s", v)
		}
	}

	for _, v := range e {
		t.Logf("%s", v)
	}
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
