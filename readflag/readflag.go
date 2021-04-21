package readflag

import (
	"flag"
)

type Flag struct {
	ConfigFile   string
	Port         uint
	Db_url       string
	Kafka_broker string
}

var ConfigFlag Flag

func init() {

	flag.StringVar(&ConfigFlag.ConfigFile, "f", "test.yaml", "config file")
	flag.UintVar(&ConfigFlag.Port, "p", 0, "port")
	flag.StringVar(&ConfigFlag.Db_url, "db", "", "db connect")
	flag.Parse()
}
