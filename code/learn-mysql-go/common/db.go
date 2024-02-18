package common

import "database/sql"

/*
Note:
1. start mysql server by `install/docker/single_node`
2. create test_db
*/

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
