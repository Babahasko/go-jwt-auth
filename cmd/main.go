package main

import (
	"fmt"
	"github.com/Babahasko/go-jwt-auth/configs"
)

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)
}
