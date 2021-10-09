package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

//Config - struct to storage the configuration from the file conf
type Config struct {
	Csv struct {
		Path string
		Name string
	}
	Server struct {
		Port string
	}
	Client struct {
		Host   string
		APIVer string
	}
}

var c Config

//ReadConfig - Reads the configuration from the file and storage it on the struct
func ReadConfig() Config {
	Config := &c

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return *Config
}
