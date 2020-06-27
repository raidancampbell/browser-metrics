package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

const configFilename = "config.yml"

type conf struct {
	ListenPort         int
	LogLevel           string
	DatasourceLocation string
}

func Initialize() *conf {
	var c conf
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFilename)
	err := viper.ReadInConfig()
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Failed to read in configuration from %s, '%s'", configFilename, err.Error())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(_ fsnotify.Event) {
		err := viper.Unmarshal(&c)
		if err != nil {
			log.Fatalf("Failed to update configuration from %s, '%s'", configFilename, err.Error())
		}
	})

	initLogger(&c)
	return &c
}

func initLogger(config *conf) {
	lvl, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logrus.Warnf("Unable to parse log level '%s', defaulting to info...", config.LogLevel)
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(lvl)
	}
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
}
