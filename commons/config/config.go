package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfgInstance *viper.Viper

func InitConfigReader() error {
	cfgInstance = viper.New()
	cfgInstance.AutomaticEnv()
	path := "./config/"
	cfgInstance.AddConfigPath(path)
	configFile := cfgInstance.GetString("ENV")
	cfgInstance.SetConfigName(configFile)
	cfgInstance.SetConfigType("json")

	err := cfgInstance.ReadInConfig()
	if err != nil {
		return fmt.Errorf("config read error: %w", err)
	}

	return nil
}

func GetConfig(key string) string {
	return cfgInstance.GetString(key)
}

func ReadConfigInto(config interface{}) (err error) {
	// setup path for config, filename, type
	viper.AutomaticEnv()
	path := "./config/"
	viper.AddConfigPath(path)
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
