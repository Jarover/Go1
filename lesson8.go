package main

import (
	"Go1/readconf"
	"Go1/readflag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Parsing YAML file")
	fmt.Println(readflag.ConfigFlag.ConfigFile)

	x, err := readconf.ReadConfig(readflag.ConfigFlag.ConfigFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Замена ключей конфигурации принятыми параметрами

	if readflag.ConfigFlag.Port > 0 {
		x.Port = readflag.ConfigFlag.Port
	}

	if len(readflag.ConfigFlag.Db_url) > 0 {
		err = x.CheckUrl(readflag.ConfigFlag.Db_url)
		if err == nil {
			x.Db_url = readflag.ConfigFlag.Db_url
		} else {
			fmt.Println(err)
			os.Exit(1)

		}
	}

	fmt.Println(x)

}
