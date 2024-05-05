package main

import (
	"tgwp/configs"
	"time"
)

func main() {
	configs.Init()
	for {
		print(1)
		time.Sleep(10 * time.Second)
	}
}
