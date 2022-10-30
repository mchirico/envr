package configuration

import (
	"fmt"
	"github.com/mchirico/envr/config"
	"github.com/mchirico/envr/logger"
	"os"
	"path/filepath"
)

var p *Configure

func init() {
	p = New()
}

type Configure struct {
	path string
}

func New() *Configure {
	p := &Configure{
		path: filepath.Join(HomeDirectory(), ".envr/profiles.yaml"),
	}
	config.SetPath(p.path)
	return p
}
func SetPath(path string) error {
	return p.setPath(path)
}
func (p *Configure) setPath(path string) error {
	p.path = path
	return config.SetPath(p.path)

}

func (p *Configure) log(msg string) {
	logger.Log(msg)
}

func PR(key string) (string, error) {
	out := ""
	env, err := ProfileEnvExports(key)
	if err != nil {
		profiles := profilesAvailable()
		fmt.Println("Configure not found:", key)
		fmt.Println("Available profiles are:")
		for _, p := range profiles {
			fmt.Println(p)
		}
		os.Exit(1)

	}
	fmt.Println(env)
	return out, nil
}

func ProfilesAvailable() []string {
	return profilesAvailable()
}

func profilesAvailable() []string {
	m := config.GetMap("profiles")

	out := []string{}
	for k := range m {
		out = append(out, k)
	}
	return out
}

func ProfileEnvExports(key string) (map[string]string, error) {
	p.log("ProfileEnvExports: profiles." + key + ".env")
	return p.exports("profiles." + key + ".env")
}

func (p *Configure) exports(key string, opt ...string) (map[string]string, error) {

	m := map[string]string{}
	for _, v := range config.GetMap(key) {
		for k, v := range v.(map[string]interface{}) {
			m[k] = v.(string)
		}

	}

	p.log("exports output returned: " + key + " " + fmt.Sprintf("\n%v\n", m))
	return m, nil
}

func HomeDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %s", err)
	}
	return home
}
