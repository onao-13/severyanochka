package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbName string
}

var log = logrus.New()

func LoadDevConfig() Config {
	v := viper.New()
	v.SetDefault("PORT", "8088")
	v.SetDefault("DBUSER", "postgres")
	v.SetDefault("DBPASS", "pass")
	v.SetDefault("DBHOST", "localhost")
	v.SetDefault("DBPORT", "5433")
	v.SetDefault("DBNAME", "severyanochka")

	var config Config

	err := v.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Config pars error")
	}

	return config
}

func LoadProdConfig() Config {
	v := viper.New()
	v.SetDefault("PORT", "8088")
	v.SetDefault("DBUSER", "postgres")
	v.SetDefault("DBPASS", "postgres")
	v.SetDefault("DBHOST", "postgres")
	v.SetDefault("DBPORT", "5432")
	v.SetDefault("DBNAME", "severyanochka")

	var config Config

	err := v.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Config pars error")
	}

	return config
}

func (config *Config) DbUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", config.DbUser,
		config.DbPass, config.DbHost, config.DbPort, config.DbName,
	)
}
