package config

import (
	"github.com/mchirico/envr/fixtures"
	"testing"
)

func TestReadConfigFile(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Simple Test",
			want:    "{{TOKEN}}",
			wantErr: false,
		},
	}
	found := false
	for _, tt := range tests {
		SetPath(fixtures.Path(".envr"))
		t.Run(tt.name, func(t *testing.T) {
			i := GetMap("profiles.argo.env")
			for k, v := range i {
				if k == "token" {
					if result, ok := v.(map[string]interface{})["replace_string"]; ok {
						if result == tt.want {
							found = true
						}
					}

				}
			}

		})
	}
	if !found {
		t.Errorf("ReadConfigFile() = %v", found)
	}
}

func TestGetMap(t *testing.T) {
	SetPath(fixtures.Path(".envr"))
	m := GetMap("profiles.argo.env")
	for k, v := range m {
		t.Logf("%s: %s", k, v)
	}
}

func TestLogger(t *testing.T) {
	SetPath(fixtures.Path(".envr"))
	m := Get("log.maxsize")
	t.Logf("%s", m)
}

func TestGetProfiles(t *testing.T) {
	SetPath(fixtures.Path(".envr"))
	m := GetMap("profiles")
	for k, _ := range m {
		t.Logf("%s", k)
	}
}
