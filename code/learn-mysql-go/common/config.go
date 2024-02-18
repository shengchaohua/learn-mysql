package common

import "time"

type Config struct {
	// common
	Host     string
	Port     string
	User     string
	Password string
	DBName   string

	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifeTime      time.Duration
	ShowSQL              bool
	AllowNativePasswords bool
}
