package main

import (
	"log"
	"time"

	"learn-mysql-go/common"
)

func TestMaxIdleConns() {
	log.Println("TestMaxIdleConns")

	db := common.OpenDB()
	db.SetMaxOpenConns(0) // no limit
	db.SetMaxIdleConns(1)
	db.SetConnMaxIdleTime(time.Second * 3) // for test

	for i := 0; i < 10; i++ {
		var connectionID int
		err := db.QueryRow("select CONNECTION_ID()").Scan(&connectionID)
		if err != nil {
			return
		}

		log.Println("connection id:", connectionID)
		time.Sleep(time.Second * 5)
	}

	log.Println("DONE!")
}

/*
Output:
2024/02/20 09:20:25 TestMaxIdleConns
2024/02/20 09:20:25 connection id: 90
2024/02/20 09:20:30 connection id: 92
2024/02/20 09:20:35 connection id: 93
2024/02/20 09:20:40 connection id: 94
2024/02/20 09:20:45 connection id: 95
2024/02/20 09:20:50 connection id: 96
2024/02/20 09:20:55 connection id: 97
2024/02/20 09:21:00 connection id: 99
2024/02/20 09:21:05 connection id: 100
2024/02/20 09:21:10 connection id: 101
2024/02/20 09:21:15 DONE!
*/
