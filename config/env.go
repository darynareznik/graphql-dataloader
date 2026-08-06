package config

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func Read() Config {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")

	viper.SetDefault("LOG_LEVEL", "DEBUG")
	viper.SetDefault("GRAPHQL_PORT", 8080)
	viper.SetDefault("GRAPHQL_ENABLE_PLAYGROUND", true)

	return Config{
		LogLevel:                viper.GetString("LOG_LEVEL"),
		GraphQLPort:             viper.GetInt("GRAPHQL_PORT"),
		GraphQLEnablePlayground: viper.GetBool("GRAPHQL_ENABLE_PLAYGROUND"),
	}
}

func (c Config) GetLogLevel() zerolog.Level {
	l, err := zerolog.ParseLevel(c.LogLevel)
	if err != nil {
		return zerolog.InfoLevel
	}

	return l
}
