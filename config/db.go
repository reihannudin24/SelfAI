package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"log"
	"time"
)

func Connection() *sql.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/self_ai?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	log.Println("Database connection successfully")
	return db
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		if errorRollback != nil {
			panic(errorRollback)
		}
		panic(err)
	} else {
		errorCommit := tx.Commit()
		if errorCommit != nil {
			panic(errorCommit)
		}
	}
}
