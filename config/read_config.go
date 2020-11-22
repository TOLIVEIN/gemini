package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//Config ...
type Config struct {
	Port     string
	RunMode  string
	PageSize string
	// Mysql MySQL
	Mysql struct {
		URL      string
		Username string
		Password string
		Database string
	}
}

//MySQL ...
// type MySQL struct {
// 	url      string
// 	username string
// 	password string
// 	database string
// }

//CONFIG ...
var CONFIG Config

//ReadConfig ...
func ReadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("read config file error: %s\n", err)
		os.Exit(1)
	}

	// viper.WatchConfig()

	err = viper.Unmarshal(&CONFIG)
	if err != nil {
		fmt.Printf("fail to unmarshal config: %s\n", err)
	}
	fmt.Println(CONFIG)
}
