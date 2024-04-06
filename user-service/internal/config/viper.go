package config

import (
	"go-microservice/user-service/internal/common"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func LoadConfig() {
	viper := viper.New()

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()

	common.PanicIfError(err)

	Config = viper
}
