package fixtures

import (
	"os"
	"path/filepath"
	"runtime"
)

// basepath is the root directory of this package.
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path,
// relative to the "fixtures" directory.
// If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}

func Read(rel string) ([]byte, error) {
	return os.ReadFile(Path(rel))
}

func Write(rel string, data []byte) error {
	return os.WriteFile(Path(rel), data, 0644)
}
