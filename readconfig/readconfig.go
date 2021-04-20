package readconfig

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config - структура для считывания конфигурационного файла
type Config struct {
	Db_url       string `yaml:"db_url"`
	Port         int    `yaml:"port"`
	Jaeger_url   string `yaml:"jaeger_url"`
	Sentry_url   string `yaml:"sentry_url"`
	Kafka_broker string `yaml:"kafka_broker"`
	Some_app_id  string `yaml:"some_app_id"`
	Some_app_key string `yaml:"some_app_key"`
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

	return x, nil
}
