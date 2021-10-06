package config

import (
	"fmt"
	"log"
	"os"

	//"github.com/davecgh/go-spew/spew"

	"github.com/spf13/viper"
)

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

var C Config

func ReadConfig() Config {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	//viper.AddConfigPath(filepath.Join("$GOPATH", "src", "github.com", "s1nuh3", "academy-go-q32021", "config"))
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
	//spew.Dump(C)
}
