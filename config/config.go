package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var GlobalConfig Application

type Application struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint16 `yaml:"port"`
	}
	Database struct {
		Host     string `yaml:"host"`
		Dbname   string `yaml:"dbname"`
		Port     uint16 `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		Db       uint8  `yaml:"db"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

func init() {
	log.Println("start read config......")

	configFilePath := "./config/application.yaml"

	// 读取配置文件
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Printf("config init failure,err: %s\n", err)
		log.Fatal("Application stop")
	}

	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		log.Printf("config file read failure,err: %s\n", err)
		log.Fatal("Application stop")
	}
	if n == 0 {
		log.Println("read config data lengh is 0")
		log.Fatal("Application stop")
	}

	application := Application{}
	err = yaml.Unmarshal(data[0:n], &application)
	if err != nil {
		log.Printf("analysis config failure,err: %s\n", err)
		log.Fatal("Application stop")
	}
	GlobalConfig = application
}
