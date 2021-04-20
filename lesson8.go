package main

import (
	rc "Go1/readconfig"
	"Go1/readflag"
	"fmt"
)

func main() {
	fmt.Println("Parsing YAML file")
	fmt.Println(readflag.ConfigFlag.ConfigFile)

	x, _ := rc.ReadConfig(readflag.ConfigFlag.ConfigFile)

	fmt.Println(x)

}
