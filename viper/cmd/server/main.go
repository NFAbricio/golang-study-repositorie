package main

import (
	"fmt"
	"log"

	"viper/config"
)

func main() {
	env, err := config.LoadConf("../../")
	if err != nil {
		log.Fatalf("error to load:%v", err)
	}
	fmt.Println(env.DatabaseName)
}
