package config

import "github.com/spf13/viper"

type Config struct {
	UsersServiceAddress string `mapstructure:"USERS_SERVICE_ADDRESS"`
	PostsServiceAddress string `mapstructure:"POSTS_SERVICE_ADDRESS"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	TokenKey            string `mapstructure:"TOKEN_KEY"`
	Environment         string `mapstructure:"ENVIRONMENT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
