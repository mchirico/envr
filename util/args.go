package util

import (
	"fmt"
	"os"
	"strings"
)

func Get(key string) (string, error) {
	env := os.Getenv(key)
	if env == "" {
		return "", fmt.Errorf("key %s not set", key)
	}
	return env, nil

}

func GetAWS() ([]string, error) {
	if _, err := Get("AWS_DEFAULT_REGION"); err != nil {
		_ = os.Setenv("AWS_DEFAULT_REGION", "us-east-1")
	}

	out := []string{}
	for _, key := range []string{"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
		"AWS_DEFAULT_REGION"} {
		env := os.Getenv(key)
		if env == "" {
			return nil, fmt.Errorf("key %s not set", key)
		}
		out = append(out, env)
	}
	return out, nil
}

func ReadReplace(file string, out []string) (string, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	s := strings.Replace(string(dat), "${AWS_ACCESS_KEY_ID}", fmt.Sprintf("%q", out[0]), -1)
	s = strings.Replace(s, "${AWS_SECRET_ACCESS_KEY}", fmt.Sprintf("%q", out[1]), -1)
	s = strings.Replace(s, "${AWS_SESSION_TOKEN}", fmt.Sprintf("%q", out[2]), -1)
	s = strings.Replace(s, "${AWS_DEFAULT_REGION}", fmt.Sprintf("%q", out[3]), -1)

	return s, nil
}
