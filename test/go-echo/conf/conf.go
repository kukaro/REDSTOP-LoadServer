package conf

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
)

const (
	BINDATA = "DEFAULT"
)

var (
	Conf              config
	defaultConfigFile = "conf/conf.toml"
)

type config struct {
	App      app
	Database database
	Static   static
}

type app struct {
	Name string `toml:"name"`
}

type static struct {
	Type string `toml:"type"`
}

type database struct {
	Name     string
	UserName string
	Password string
	Host     string
	Port     string
}

func Init(configFile string) error {
	if configFile == "" {
		configFile = defaultConfigFile
	}
	Conf = config{}
	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file err : " + err.Error())
	} else {
		log.Infof("load config from file : " + configFile)
		configBytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			return errors.New("config load err : " + err.Error())
		}
		_, err = toml.Decode(string(configBytes), &Conf)
		if err != nil {
			return errors.New("config decode err : " + err.Error())
		}
	}
	log.Infof("config data:%v", Conf)
	return nil
}
