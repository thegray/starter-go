package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type appConfig struct {
	Server serverConfig `yaml:"server" mapstructure:"server"`
	Logger loggerConfig `yaml:"logger" mapstructure:"logger"`
	DB     databaseConfig `yaml:"db" mapstructure:"db"`
}

var cfg = new(appConfig)

func Init(path string) error {
	fmt.Printf("reading config path: %s\n", path)
	
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	// Set a prefix for environment variables, so we can override config with APP_ prefixed env vars.
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	// a.b.c will be APP_A_B_C
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file, %s", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unable to decode into struct, %s", err)
	}

	return nil
}
