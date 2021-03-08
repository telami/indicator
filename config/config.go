package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var Instance *Config

type Config struct {
	Cron string `yaml:"cron"`

	// server酱
	ServerChan struct {
		Secret string `yaml:"secret"`
	} `yaml:"serverChan"`

	// 指标申请人
	Applicant []struct {
		Name     string `yaml:"name"`
		Nickname string `yaml:"nickname"`
		Code     string `yaml:"code"`
	} `yaml:"applicant"`
}

func Init(filename string) *Config {
	Instance = &Config{}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		logrus.Error(err)
	} else if err = yaml.Unmarshal(yamlFile, Instance); err != nil {
		logrus.Error(err)
	}
	return Instance
}
