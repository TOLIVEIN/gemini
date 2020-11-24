package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/viper"
)

//Config ...
type Config struct {
	Port          string
	RunMode       string
	PageSize      string
	JwtSecret     string
	RootPath      string
	PathSeparator string
	// Mysql MySQL
	Mysql struct {
		URL      string
		Username string
		Password string
		Database string
	}
	File struct {
		ImagePrefixURL string
		ImageSavePath  string
		ImageMaxSize   string
		ImageAllowExts string
	}
}

//MySQL ...
// type MySQL struct {
// 	url      string
// 	username string
// 	password string
// 	database string
// }

var config *Config

//Init ...
func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("read config file error: %s\n", err)
		os.Exit(1)
	}

	// viper.WatchConfig()

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("fail to unmarshal config: %s\n", err)
	}

	if runtime.GOOS == "windows" {
		config.RootPath = "\\"
		config.File.ImageSavePath = "upload\\images\\"
		config.PathSeparator = "\\"
	}
	// fmt.Println(config)
}

//GetConfig ...
func GetConfig() *Config {
	return config
}
