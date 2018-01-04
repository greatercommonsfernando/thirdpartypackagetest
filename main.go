package main

import (
	"fmt"
	"github.com/olebedev/config"
)

func main() {
	if cfg, err := config.ParseYamlFile("config.yml"); err != nil {
		panic(err)
	} else {
		if testValue, err := cfg.String("test"); err != nil {
			panic(err)
		} else {
			fmt.Println("test:", testValue)
		}
	}
	fmt.Println("Hello World!")
}
