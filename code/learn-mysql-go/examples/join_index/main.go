package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	"learn-mysql-go/common"
)

/*
Note:
1. Start mysql
2. Create tables with example.sql
3. Run main
*/

func main() {
	db := common.OpenDB()

	insert(db, 10000)
}

func insert(db *sql.DB, count int) {
	rand.Seed(1)

	categories := []string{"aaa", "bbb"}
	now := time.Now()
	for i := 0; i < count; i++ {
		randID := rand.Uint64()
		category := categories[randID%2]

		_, err := db.Exec("INSERT INTO join_index_1_tab (category, uid) VALUES (?, ?);", category, randID)
		if common.CheckDuplicate(err) {
			log.Println("WARN: duplicate", randID)
		} else {
			common.PanicIf(err)
		}

		if i < 3 {
			log.Println("INSERT data", i, randID)
		}
		if i > 0 && i%(count/10) == 0 {
			log.Println("Elapsed", time.Since(now))
		}
	}

	log.Println("Total", time.Since(now))
}
