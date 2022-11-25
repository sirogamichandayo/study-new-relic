package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	NewRelic NewRelic `yaml:newRelic`
}

type NewRelic struct {
	LicenseKey string `yaml:licenseKey`
}

var cfg Config

func GetConfig() Config {
	return cfg
}

func init() {
	viper.SetConfigName("local.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("env/")

	err := viper.ReadInConfig()
	if err != nil {
		panic("read error")
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic("unmarshal error")
	}
}
