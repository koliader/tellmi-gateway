package config

import sdkconfig "github.com/koliader/tellmi-sdk/config"

type Config struct {
	UsersServiceAddress string `mapstructure:"USERS_SERVICE_ADDRESS"`
	PostsServiceAddress string `mapstructure:"POSTS_SERVICE_ADDRESS"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	TokenKey            string `mapstructure:"TOKEN_KEY"`
	Environment         string `mapstructure:"ENVIRONMENT"`
}

func LoadConfig(path string) (Config, error) {
	var cfg Config
	err := sdkconfig.LoadConfig(path, &cfg)
	return cfg, err
}
