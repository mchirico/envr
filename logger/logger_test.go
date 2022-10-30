package logger

import (
	"github.com/mchirico/envr/config"
	"github.com/mchirico/envr/fixtures"
	"testing"
)

func TestLog(t *testing.T) {
	config.SetPath(fixtures.Path(".envr"))
	OverrideFileLoc(fixtures.Path("../.envr.log"))

	Log("Test more")
}
