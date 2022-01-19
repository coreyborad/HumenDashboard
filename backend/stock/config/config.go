package config

import (
	"errors"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type (
	config struct {
		Load       bool
		Name       string
		Usage      string
		Version    string
		TTL        time.Duration
		RefreshTTL time.Duration
		Debug      bool   `yaml:"debug"`
		Key        string `yaml:"app_key"`
		Server     server `yaml:"server"`
		// GRPC       grpc     `yaml:"grpc"`
		// Mail       mail     `yaml:"mail"`
		// MQTT       mqtt     `yaml:"mqtt"`
		// Redis      redis    `yaml:"redis"`
		Database database `yaml:"database"`
		MongoDB  mongodb  `yaml:"mongodb"`
		Telegram telegram `yaml:"telegram"`
		// Swagger    swagger  `yaml:"swagger"`
	}
	server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		// HTTPS       bool   `yaml:"https"`
		// LetsEncrypt bool   `yaml:"lets_encrypt"`
		// Cert        string `yaml:"cert"`
		// Key         string `yaml:"key"`
		// GeoIP       string `yaml:"geoip"`
	}
	// grpc struct {
	// 	Port    string  `yaml:"port"`
	// 	Account account `yaml:"account"`
	// }
	// mail struct {
	// 	Host string `yaml:"host"`
	// 	Port int    `yaml:"port"`
	// 	From struct {
	// 		Address string `yaml:"address"`
	// 		Name    string `yaml:"name"`
	// 	} `yaml:"from"`
	// 	Username   string `yaml:"username"`
	// 	Password   string `yaml:"password"`
	// 	Encryption string `yaml:"encryption"`
	// }
	// mqtt struct {
	// 	Scheme string `yaml:"scheme"`
	// 	Broker string `yaml:"broker"`
	// 	Port   string `yaml:"port"`
	// 	CA     string `yaml:"ca"`
	// 	Cert   string `yaml:"cert"`
	// 	Key    string `yaml:"key"`
	// }
	// redis struct {
	// 	Host   string `yaml:"host"`
	// 	Port   string `yaml:"port"`
	// 	Prefix string `yaml:"prefix"`
	// }
	database struct {
		Default string `yaml:"default"`
		Mysql   struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Dbname   string `yaml:"dbname"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			TLS      string `yaml:"tls"`
			Charset  string `yaml:"charset"`
		} `yaml:"mysql"`
		Pgsql struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Username string `yaml:"username"`
			Dbname   string `yaml:"dbname"`
			Password string `yaml:"password"`
		} `yaml:"pgsql"`
		Sqlite struct {
			Driver string `yaml:"driver"`
			Dbname string `yaml:"dbname"`
		} `yaml:"sqlite"`
		Sqlsrv struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Username string `yaml:"username"`
			Dbname   string `yaml:"dbname"`
			Password string `yaml:"password"`
		} `yaml:"sqlsrv"`
	}
	mongodb struct {
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		DbName string `yaml:"dbname"`
	}
	telegram struct {
		Token string   `yaml:"token"`
		Chats []string `yaml:"chats"`
	}
	// swagger struct {
	// 	URL string `yaml:"url"`
	// }
)

// default value
var (
	// App App
	App = &config{
		Load:       false,
		Name:       "ABC",
		Usage:      "AAA",
		Version:    "1.0.0",
		TTL:        3600,
		RefreshTTL: 1209600,
	}
	Server   = &server{}
	Database = &database{}
	MongoDB  = &mongodb{}
	Telegram = &telegram{}
	// GRPC     = &grpc{}
	// Mail     = &mail{}
	// MQTT     = &mqtt{}
	// Redis    = &redis{}
	// Swagger  = &swagger{}
)

// Load Load configuration
func Load(path string) (err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, App)
	if err != nil {
		return err
	}

	Server = &App.Server
	// GRPC = &App.GRPC
	// Mail = &App.Mail
	// MQTT = &App.MQTT
	// Redis = &App.Redis
	Database = &App.Database
	MongoDB = &App.MongoDB
	// Swagger = &App.Swagger
	Telegram = &App.Telegram

	App.Load = true

	return nil
}

// Check Check
func Check() (err error) {
	if !App.Load {
		err = errors.New("config does not load configuration file")
	}

	return
}
