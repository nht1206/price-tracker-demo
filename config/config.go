package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AppName                string `split_words:"true"`
	NumCrawlingGoroutines  int    `split_words:"true"`
	NumNotifyingGoroutines int    `split_words:"true"`
	Log                    *LogConfig
	DB                     *DatabaseConfig
	Notifier               *NotifierConfig
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

type NotifierConfig struct {
	Mail *MailConfig
}

type MailConfig struct {
	SMTPHost       string `envconfig:"NOTIFIER_MAIL_SMTP_HOST"`
	SMTPPort       string `envconfig:"NOTIFIER_MAIL_SMTP_PORT"`
	Sender         string
	SenderPassword string `split_words:"true"`
}
