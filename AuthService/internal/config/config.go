package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Cfg struct {
	Env          string `yaml:"env" env-default:"local"`
	Storage      string `yaml:"storage" env-default:"database"`
	StoragePort  string `yaml:"storage_port" env-default:"5432"`
	StorageLogin string `yaml:"storage_login"`
	StoragePass  string `yaml:"storage_pass"`
	GRPConfig    GRPConfig
}

type GRPConfig struct {
	Port    int           `yaml:"port" port-default:"440444"`
	Timeout time.Duration `yaml:"timeout" timeout-default:"1h"`
}

func MustLoad() *Cfg {
	configPath := fetchConfigPath()
	if configPath == "" {
		configPath = "../../configs/base.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist " + configPath)
	}

	var conf Cfg

	if err := cleanenv.ReadConfig(configPath, &conf); err != nil {
		panic("config path is empty" + err.Error())
	}

	return &conf
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
