package common

import "github.com/go-sql-driver/mysql"

func CheckDuplicate(err error) bool {
	if errMySQL, ok := err.(*mysql.MySQLError); ok {
		switch errMySQL.Number {
		case 1062:
			return true
		}
	}

	return false
}
