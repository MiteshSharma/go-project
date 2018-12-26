package setting

import (
	"fmt"
	"os"

	"github.com/MiteshSharma/project/model"
	"github.com/spf13/viper"
)

func GetConfig() *model.Config {
	return GetConfigFromFile("default")
}

func GetConfigFromFile(fileName string) *model.Config {
	if fileName == "" {
		fileName = "default"
	}
	viper.SetConfigName(fileName)
	viper.AddConfigPath("../conf/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf/")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}

	config := &model.Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("couldn't read config: %s", err)
		os.Exit(1)
	}
	return config
}
