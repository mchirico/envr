package util

import (
	"fmt"
	"github.com/mchirico/envr/pb"
	"os"
	"regexp"
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

func AWSFromPB() ([]string, error) {
	out := []string{}
	m := map[string]string{}
	order := []string{"export AWS_ACCESS_KEY_ID",
		"export AWS_SECRET_ACCESS_KEY", "export AWS_SESSION_TOKEN"}
	for _, v := range order {
		m[v] = ""
	}

	r, err := pb.Read()
	if err != nil {
		return nil, err
	}
	t := strings.Split(r, "\n")
	for _, v := range t {
		o := strings.Split(v, "=")
		if len(o) != 2 {
			return nil, fmt.Errorf("error: %s", v)
		}
		if _, ok := m[o[0]]; ok {
			m[o[0]] = strings.Replace(o[1], "\"", "", -1)
		}
	}
	for _, v := range order {
		out = append(out, m[v])
	}
	if v, err := Get("AWS_DEFAULT_REGION"); err != nil {
		out = append(out, "us-east-1")
	} else {
		out = append(out, v)
	}
	return out, nil
}

const ENVREGX = `{[A-Z_\-]+[0-9]*}`

func FindAll(file string) ([]string, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return []string{}, err
	}
	re := regexp.MustCompile(ENVREGX)

	return re.FindAllString(string(dat), -1), nil
}

type Map struct {
	Replacement  string
	Count        int
	ReplaceCount int
}

type Env struct {
	m map[string]Map
}

func NewEnv(keys []string) *Env {
	e := &Env{m: map[string]Map{}}
	e.init(keys)
	return e
}

func (e *Env) init(keys []string) {
	e.m = make(map[string]Map)
	for _, v := range keys {
		if val, ok := e.m[v]; ok {
			val.Count++
			e.m[v] = val
		} else {
			if e.getEnv(v) == "" {
				continue
			}
			e.m[v] = Map{Replacement: e.getEnv(v), Count: 1}
		}
	}
}

func (e *Env) getEnv(key string) string {
	if len(key) >= 3 {
		return os.Getenv(key[1 : len(key)-1])
	}
	return ""
}

func (e *Env) Replace(fileIn, fileOut string) error {
	s, err := os.ReadFile(fileIn)
	if err != nil {
		return err
	}
	for k, v := range e.m {
		if v.Count == 0 {
			continue
		}
		s = []byte(strings.Replace(string(s), k, v.Replacement, -1))
		v.ReplaceCount++
		v.Count--
		e.m[k] = v
	}
	return os.WriteFile(fileOut, s, 0644)
}
