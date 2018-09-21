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
	Server   server
	TestServer testServer
}

type app struct {
	Name  string `toml:"name"`
	Owner string `toml:"owner"`
}

type static struct {
	Type string `toml:"type"`
}

type server struct {
	Addr            string `toml:"addr"`
	DomainApi       string `toml:"domain_api"`
	DomainWeb       string `toml:"domain_web"`
	DomainWebSocket string `toml:"domain_web_socket"`
}

type database struct {
	Name     string
	UserName string
	Password string
	Host     string
	Port     string
}

type testServer struct {
	TestDomain string `toml:"test_domain"`
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
