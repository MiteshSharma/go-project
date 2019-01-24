package setting

import (
	"fmt"
	"os"
	"sync"

	"github.com/MiteshSharma/project/model"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var watchOnce sync.Once
var mutexConfigListener sync.Mutex
var mapConfigListener = map[IConfigChangeListener]struct{}{}

type IConfigChangeListener interface {
	OnConfigChange(newConfig *model.Config)
}

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

	return getUpdatedConfig()
}

func getUpdatedConfig() *model.Config {
	config := &model.Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("couldn't read config: %s", err)
		os.Exit(1)
	}
	return config
}

func WatcherConfig() {
	watchOnce.Do(func() {
		viper.WatchConfig()
		viper.OnConfigChange(onConfigChange)
	})
}

func onConfigChange(event fsnotify.Event) {
	mutexConfigListener.Lock()
	defer mutexConfigListener.Unlock()
	newConfig := getUpdatedConfig()
	for listener := range mapConfigListener {
		listener.OnConfigChange(newConfig)
	}
}

func AddConfigChangeListener(listener IConfigChangeListener) {
	mutexConfigListener.Lock()
	defer mutexConfigListener.Unlock()
	mapConfigListener[listener] = struct{}{}
}

func DeleteConfigChangeListener(listener IConfigChangeListener) {
	mutexConfigListener.Lock()
	defer mutexConfigListener.Unlock()
	delete(mapConfigListener, listener)
}
