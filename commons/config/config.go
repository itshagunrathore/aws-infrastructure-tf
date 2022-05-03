package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfigInto(config interface{}) (err error) {
	path := "./config/"
	viper.AddConfigPath(path)

	viper.AutomaticEnv()
	configFile := viper.GetString("ENV")
	viper.SetConfigName(configFile)
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("config read error: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return fmt.Errorf("config unmarshalling error: %w", err)
	}

	return nil
}
