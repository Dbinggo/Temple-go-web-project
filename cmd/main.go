package main

import (
	"tgwp/configs"
	"tgwp/internal/db/mySql"
	"tgwp/log"
)

func main() {
	log.InitLogger()
	configs.Init()
	mySql.InitMySql()

}
