package config

import "github.com/spf13/viper"

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"localhost"`
		Port   string `yaml:"port" env-default:"8080"`
	}

	PostgresDB struct {
		Host     string `yaml:"host" env-required:"true"`
		Port     string `yaml:"port" env-required:"true"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database" env-required:"true"`
	}
}

func LoadConfig(path string) (Config, error) {
	config := Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
