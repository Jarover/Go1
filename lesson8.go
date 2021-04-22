package main

import (
	"Go1/readconf"
	"Go1/readflag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Parsing YAML config file - ", readflag.ConfigFlag.ConfigFile)

	x, err := readconf.ReadConfig(readflag.ConfigFlag.ConfigFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Замена ключей конфигурации принятыми параметрами

	if readflag.ConfigFlag.Port > 0 {
		x.SetPort(readflag.ConfigFlag.Port)
	}

	if len(readflag.ConfigFlag.Db_url) > 0 {
		x.SetDb(readflag.ConfigFlag.Db_url)
	}
	// проверяем результирующий конфиг
	fmt.Println(x)

}
