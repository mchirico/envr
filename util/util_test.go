package util

import (
	"github.com/mchirico/envr/fixtures"
	"github.com/mchirico/envr/pb"
	"os"
	"testing"
)

// live test
func TestGetAWS(t *testing.T) {

	r, err := GetAWS()
	if err != nil {
		return
	}
	_ = r
}

func writeMockCreds() {
	pb.Write(`export AWS_ACCESS_KEY_ID="ID"
export AWS_SECRET_ACCESS_KEY="secret"
export AWS_SESSION_TOKEN="token"`)

}

func TestAWSFromPB(t *testing.T) {
	writeMockCreds()
	r, err := AWSFromPB()
	if err != nil {
		t.Error(err)
	}

	if len(r) != 4 {
		t.Error("Should be 4")
	}
	if r[2] != "token" {
		t.Error("Should be token")
	}

}

func TestFindAll(t *testing.T) {
	os.Setenv("WITH2", "2")
	file := fixtures.Path("./findall/sampleFile.yaml")
	r, _ := FindAll(file)
	e := NewEnv(r)

	fileOut := fixtures.Path("./findall/sampleFileOut.yaml")
	e.Replace(file, fileOut)

	// Check for ${WITH2} replaced
	r, _ = FindAll(fileOut)
	if r[0] != "{WITH1}" && r[1] != "{WITH1}" && len(r) != 2 {
		t.Error("Should be {WITH1} and {WITH1}")
	}

}
