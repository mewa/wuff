package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/mewa/wuff/config"
)

var conf config.Config

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.SetConfigType("hcl")

	conf, err := config.ReadConfig()

	if err != nil {
		log.Fatalln(err)
	}

	err = conf.Verify()

	if err != nil {
		log.Fatalln(err)
	}

	log.Info("Loaded config")
	for _, service := range conf.Service {
		log.WithFields(log.Fields{
			"name": service.Name,
			"checkPeriod": service.CheckPeriod,
			"retries": service.Retries,
			"retryPeriod": service.RetryPeriod,
			"check": service.Check,
		}).Infof("Loaded service: %s", service.Name)
	}
}
