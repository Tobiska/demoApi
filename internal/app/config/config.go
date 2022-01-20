package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	IsDebug bool `yaml:"IS_DEBUG"`
	Listen  struct {
		Type   string `mapstructure:"TYPE"`
		BindIP string `mapstructure:"BIND_IP"`
		Port   string `mapstructure:"PORT"`
	}

	PostgresDB struct {
		Host     string `mapstructure:"DB_HOST, env-default="`
		Port     string `mapstructure:"DB_PORT"`
		Username string `mapstructure:"DB_USERNAME"`
		Password string `mapstructure:"DB_PASSWORD"`
		Database string `mapstructure:"DB_DATABASE"`
	}
}

func LoadConfig(path string) (Config, error) {
	//TODO add sync once
	vi := viper.New()
	config := Config{}
	vi.SetConfigFile(path)
	vi.SetConfigType("env")

	vi.AutomaticEnv()
	if err := vi.ReadInConfig(); err != nil {
		return config, err
	}
	fmt.Println(vi.AllSettings())
	err := vi.Unmarshal(&config.PostgresDB)
	err = vi.Unmarshal(&config.Listen)
	err = vi.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
