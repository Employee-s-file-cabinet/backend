package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbOpenError = "Open DataBase Error"
)

func OpenDB(DBDSN string) (*sql.DB, error) {
	db, errDB := sql.Open("postgres", DBDSN)
	log.Println(DBDSN)
	if errDB != nil {
		log.Println(dbOpenError)
		log.Println(errDB)
	}

	// Проверка подключения
	errDB = db.Ping()
	if errDB != nil {
		log.Println(errDB)
	}
	return db, errDB
}
