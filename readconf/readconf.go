package readconf

import (
	"io/ioutil"
	"net/url"

	"gopkg.in/yaml.v2"
)

// Config - структура для считывания конфигурационного файла
type Config struct {
	Db_url       string `yaml:"db_url"`
	Port         uint   `yaml:"port"`
	Jaeger_url   string `yaml:"jaeger_url"`
	Sentry_url   string `yaml:"sentry_url"`
	Kafka_broker string `yaml:"kafka_broker"`
	Some_app_id  string `yaml:"some_app_id"`
	Some_app_key string `yaml:"some_app_key"`
}

func (e *Config) Validate() error {
	var err error
	err = e.CheckUrl(e.Db_url)
	if err != nil {
		return err
	}

	err = e.CheckUrl(e.Jaeger_url)
	if err != nil {
		return err
	}

	err = e.CheckUrl(e.Sentry_url)
	if err != nil {
		return err
	}

	err = e.CheckUrl(e.Kafka_broker)
	if err != nil {
		return err
	}

	return nil
}

func (e *Config) SetPort(p uint) error {

	e.Port = p
	return nil

}

func (e *Config) SetDb(p string) error {

	err := e.CheckUrl(p)
	if err == nil {
		e.Db_url = p
		return nil
	} else {
		return err

	}

}

func (e *Config) CheckUrl(path string) error {

	_, err := url.ParseRequestURI(path)

	if err != nil {
		return err
	}
	return nil

}

func ReadConfig(ConfigName string) (x *Config, err error) {
	var file []byte
	if file, err = ioutil.ReadFile(ConfigName); err != nil {
		return nil, err
	}
	x = new(Config)
	if err = yaml.Unmarshal(file, x); err != nil {
		return nil, err
	}

	if err = x.Validate(); err != nil {
		return nil, err
	}

	return x, nil
}
