package main

import (
	"learn-mysql-go/common"
	"log"
)

func main() {
	db := common.Connect()
	log.Println("DB:", db)
}
