package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

var (
	confPath string
	conf     Cfg
)

type Cfg struct {
	GatewayPort int         `yaml:"port" port-default:"8080"`
	Auth        AuthConfig  `yaml:"authService"`
	Chat        ChatConfig  `yaml:"chatService"`
	Group       GroupConfig `yaml:"groupService"`
}

type AuthConfig struct {
	Port    int           `yaml:"port" port-default:"40444"`
	Timeout time.Duration `yaml:"timeout" timeout-default:"1h"`
}

type ChatConfig struct {
	Port    int           `yaml:"port" port-default:"440445"`
	Timeout time.Duration `yaml:"timeout" timeout-default:"1h"`
}

type GroupConfig struct {
	Port    int           `yaml:"port" port-default:"44046"`
	Timeout time.Duration `yaml:"timeout" timeout-default:"1h"`
}

func init() {
	flag.StringVar(&confPath, "config", "", "path to config")
	flag.Parse()

	if confPath == "" {
		confPath = os.Getenv("CHAT_CONFIG_PATH")
	}

	if confPath == "" {
		confPath = "./config/base.yaml"
	}
}

func Load() Cfg {
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		log.Panicf("config path is empty: %s", confPath)
	}

	if err := cleanenv.ReadConfig(confPath, &conf); err != nil {
		log.Panicf("unexpected err: %v", err)
	}

	return conf
}
