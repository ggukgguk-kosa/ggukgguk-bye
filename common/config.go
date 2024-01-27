package common

import (
	"fmt"

	"github.com/spf13/viper"
)

var config Config

func InitConfig() Config {
	viper.SetDefault("NEXT_STEP", 1)
	viper.SetConfigName("ggukgguk-bye")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config = Config{
		DBUri:    viper.GetString("DB_URI"),
		NextStep: viper.GetInt("NEXT_STEP"),
	}

	return config
}

func getConfig() Config {
	return config
}

type Config struct {
	DBUri    string
	NextStep int
}
