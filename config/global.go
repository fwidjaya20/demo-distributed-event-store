package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
	_ "github.com/lib/pq"
)

var db *sqlx.DB
var once sync.Once

// GetDb is a function to get DB instance
func GetDb() *sqlx.DB {
	once.Do(func() {
		var err error
		conn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			GetEnv(DB_HOST),
			GetEnv(DB_PORT),
			GetEnv(DB_USER),
			GetEnv(DB_PASS),
			GetEnv(DB_NAME),
		)

		db, err = sqlx.Connect("postgres", conn)
		if nil != err {
			log.Fatal(err)
		}
	})
	return db
}
