package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	PortServer    string `mapstructure:"PORT_SERVER"`
	DBString      string `mapstructure:"DB_STRING"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func ReadConfigFile(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	log.Debugf("reading config from %s/app.env", path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot read config file: %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("cannot unmarshal config file : %s", err)
	}
	return
}
