package common

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var config Config

func InitConfig() Config {
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
		FilePath: viper.GetString("FILE_PATH"),
	}

	log.Println("init config: ")
	log.Println(config)

	return config
}

func getConfig() Config {
	return config
}

type Config struct {
	DBUri    string
	NextStep int
	FilePath string
}
