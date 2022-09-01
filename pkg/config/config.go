package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type app struct {
	Version  string `yaml:"version"`
	Env      string `yaml:"env"`
	Name     string `yaml:"name"`
	Desc     string `yaml:"desc"`
	Keywords string `yaml:"keywords"`
}

type session struct {
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
}

type db struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
	Pass string `yaml:"pass"`
	DB   string `yaml:"db"`
}

type redis struct {
	Host        string `yaml:"host"`
	Pass        string `yaml:"pass"`
	Port        string `yaml:"port"`
	DB          string `yaml:"db"`
	IdleTimeout string `yaml:"idleTimeout"`
}

type upload struct {
	Path     string   `yaml:"path"`
	ImageExt []string `yaml:"imageExt"`
}

type conf struct {
	App     app     `yaml:"app"`
	Session session `yaml:"session"`
	DB      db      `yaml:"db"`
	Redis   redis   `yaml:"redis"`
	Upload  upload  `yaml:"upload"`
}

var Conf *conf

func init() {
	b, err := os.ReadFile("../config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var c *conf
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}

	Conf = c
}
