package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AppName                string `split_words:"true"`
	NumCrawlingGoroutines  int    `split_words:"true"`
	NumNotifyingGoroutines int    `split_words:"true"`
	Log                    *LogConfig
	DB                     *DatabaseConfig
}

func NewConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

type DatabaseConfig struct {
	DSN string
}

type LogConfig struct {
	OutputPath string `split_words:"true"`
	FileName   string `split_words:"true"`
	Level      string
}
