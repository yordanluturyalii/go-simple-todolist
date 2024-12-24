package config

import "github.com/spf13/viper"

type Config struct {
	Env *viper.Viper
}

func NewConfig() *Config {
	cnf := viper.New()

	cnf.SetConfigFile(".env")
	cnf.AddConfigPath(".")
	cnf.AutomaticEnv()

	err := cnf.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &Config{cnf}
}
