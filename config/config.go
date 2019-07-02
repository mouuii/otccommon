package config

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Init init config from relative path,does not include extension
func Init(path string, fileName string, config interface{}) {
	if path == "" {
		log.Fatal("config path is not given")
	}
	configPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("config abs path err: %v", err)
	}

	viper.SetConfigName(fileName)
	viper.AddConfigPath(".")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file ,%s", err)
	}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

// HotReload hotreload config
func HotReload(config interface{}) {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		reloadConfigToObject(config)
	})
	log.Printf("reload config %+v", config)
}
func reloadConfigToObject(config interface{}) {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file ,%s", err)
	}
	err := viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
