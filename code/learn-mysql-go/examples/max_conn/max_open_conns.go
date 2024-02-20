package main

import (
	"database/sql"
	"log"
	"sync"

	"learn-mysql-go/common"
)

func TestMaxOpenConns() {
	log.Println("TestMaxOpenConns")

	db := common.OpenDB()
	db.SetMaxOpenConns(1)
	var wg sync.WaitGroup

	// 开启 10 个 goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(db, i, &wg)
	}

	wg.Wait()
	log.Println("DONE!")
}

func worker(db *sql.DB, i int, wg *sync.WaitGroup) {
	defer wg.Done()

	var connectionID int
	err := db.QueryRow("SELECT CONNECTION_ID()").Scan(&connectionID)
	if err != nil {
		panic(err)
	}

	log.Printf("worker-%d uses connection id: %d", i, connectionID)

	// The workers that do not obtain mysql connection will be blocked since the MaxOpenConn is only 1.
	var result int
	err = db.QueryRow("SELECT SLEEP(5)").Scan(&result)
	if err != nil {
		panic(err)
	}

	log.Printf("worker-%d wakes up after sleeping", i)
}

/*
Output:
2024/02/20 09:21:45 TestMaxOpenConns
2024/02/20 09:21:45 worker-9 uses connection id: 103
2024/02/20 09:21:45 worker-0 uses connection id: 103
2024/02/20 09:21:45 worker-7 uses connection id: 103
2024/02/20 09:21:45 worker-1 uses connection id: 103
2024/02/20 09:21:50 worker-7 wakes up after sleeping
2024/02/20 09:21:50 worker-2 uses connection id: 103
2024/02/20 09:21:50 worker-8 uses connection id: 103
2024/02/20 09:21:55 worker-1 wakes up after sleeping
2024/02/20 09:21:55 worker-4 uses connection id: 103
2024/02/20 09:22:00 worker-8 wakes up after sleeping
2024/02/20 09:22:05 worker-9 wakes up after sleeping
2024/02/20 09:22:10 worker-0 wakes up after sleeping
2024/02/20 09:22:10 worker-5 uses connection id: 103
2024/02/20 09:22:15 worker-4 wakes up after sleeping
2024/02/20 09:22:20 worker-2 wakes up after sleeping
2024/02/20 09:22:20 worker-6 uses connection id: 103
2024/02/20 09:22:25 worker-5 wakes up after sleeping
2024/02/20 09:22:30 worker-6 wakes up after sleeping
2024/02/20 09:22:30 worker-3 uses connection id: 103
2024/02/20 09:22:35 worker-3 wakes up after sleeping
2024/02/20 09:22:35 DONE!
*/
