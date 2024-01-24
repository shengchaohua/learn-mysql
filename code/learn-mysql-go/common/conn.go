package common

import "database/sql"

func Connect() {
	driver := sql.Drivers()
	sql.OpenDB()
	db := sql.OpenDB(driver)
}
