package file

import (
	"os"
	"strings"
)

func Read(file string) ([]byte, error) {
	return os.ReadFile(file)
}

func Write(file string, data []byte) error {
	return os.WriteFile(file, data, 0644)
}

func Exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func Remove(file string) error {
	return os.Remove(file)
}

func RemoveAll(dir string) error {
	return os.RemoveAll(dir)
}

func MkdirAll(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func ReplaceFile(file string, data []byte) error {
	if err := Remove(file); err != nil {
		return err
	}
	return Write(file, data)
}

func ReplaceFileIfExist(file string, data []byte) error {
	if Exists(file) {
		return ReplaceFile(file, data)
	}
	return Write(file, data)
}

func StringReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}
