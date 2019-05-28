package env

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Address string
	Port    int
}

var Settings Config

func init() {
	fmt.Println("env")
	viper.AutomaticEnv()

	viper.SetEnvPrefix("ROOMBA")
	viper.SetDefault("ADDRESS", "0.0.0.0:")
	viper.SetDefault("PORT", 8080)

	Settings = Config{
		Address: viper.GetString("ADDRESS"),
		Port:    viper.GetInt("PORT"),
	}

}
