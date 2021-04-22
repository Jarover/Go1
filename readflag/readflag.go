package readflag

import (
	"flag"
	"os"
	"strconv"
)

type Flag struct {
	ConfigFile   string
	Port         uint
	Db_url       string
	Kafka_broker string
}

var ConfigFlag Flag

func getEnv(key string, defVal string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}

	return defVal

}

func getEnvInt(key string, defVal int64) int64 {
	if env, ok := os.LookupEnv(key); ok {
		envInt, err := strconv.ParseInt(env, 10, 64)
		if err == nil {
			return envInt
		}
	}

	return defVal

}

func init() {

	flag.StringVar(&ConfigFlag.ConfigFile, "f", getEnv("CONFIGFILE", "test.yaml"), "config file")
	flag.UintVar(&ConfigFlag.Port, "p", uint(getEnvInt("PORT", 0)), "port")
	flag.StringVar(&ConfigFlag.Db_url, "db", getEnv("DB", ""), "db connect")
	flag.Parse()
}
