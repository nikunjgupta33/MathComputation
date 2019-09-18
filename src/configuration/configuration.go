package configuration

import (
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type readConfig struct {
	ConfigFile string
}

type Configuration struct {
	Config readConfig
}

func (this *Configuration) Init() {
	fmt.Println("Init Configuration...")

	if this.Config.ConfigFile != "" {
		log.Println("Using Config File: " + this.Config.ConfigFile)
		viper.SetConfigFile(this.Config.ConfigFile)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	switch err.(type) {
	case viper.UnsupportedConfigError:
		log.Println("No Config File..!!")
	default:
		check(err)
	}

	err = viper.Unmarshal(&this.Config)
	check(err)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config File Changed: " + e.Name)
		err = viper.Unmarshal(&this.Config)
	})

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
