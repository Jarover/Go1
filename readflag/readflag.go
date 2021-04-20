package readflag

import (
	"flag"
)

type Flag struct {
	ConfigFile   string
	Port         int
	Db_url       string
	Kafka_broker string
}

var ConfigFlag Flag

func init() {

	flag.StringVar(&ConfigFlag.ConfigFile, "f", "test.yaml", "config file")
	flag.IntVar(&ConfigFlag.Port, "p", 0, "port")
	flag.StringVar(&ConfigFlag.Db_url, "db", "test.yaml", "db connect")
	flag.Parse()
}
