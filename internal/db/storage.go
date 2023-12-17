package storage

import (
	"database/sql"
	"log"

	"github.com/go-eden/slf4go"

	_ "github.com/lib/pq"
)

const (
	dbOpenError = "Open DataBase Error"
)

func OpenDB(DBDSN string) (*sql.DB, error) {

	logger := slf4go.Get("example")

	db, errDB := sql.Open("postgres", DBDSN)
	log.Println(DBDSN)
	if errDB != nil {
		log.Println(dbOpenError)
		logger.Error(err)
	}

	// Проверка подключения
	errDB = db.Ping()
	if errDB != nil {
		logger.Error(err)
	}
	return db, errDB
}
