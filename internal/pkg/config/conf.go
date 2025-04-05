package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type appConfig struct {
	Server serverConfig `yaml:"server"`
	Logger loggerConfig `yaml:"logger"`
}

var cfg = new(appConfig)

func Init(path string) error {
	fmt.Printf("reading config path: %s\n", path)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading config file, %s", err)
	}

	if err := yaml.Unmarshal(b, cfg); err != nil {
		return fmt.Errorf("unable to decode into struct, %s", err)
	}

	return nil
}
