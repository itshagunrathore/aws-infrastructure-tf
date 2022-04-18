package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

func GetConfig(config models.Configurations) (nil, err error) {
	path := "./config/"

	viper.AutomaticEnv()
	fmt.Println("ENV flag: ", os.Getenv("ENV"))
	configFile := viper.GetString("ENV")
	fmt.Println("configFile from viper: ", configFile)
	viper.AddConfigPath(path)
	viper.SetConfigName(configFile)
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		// return nil, err
		fmt.Println("some error: ",err)
	}

	err = viper.Unmarshal(&config)
	fmt.Println("error: ", err)
	return nil, nil
}
