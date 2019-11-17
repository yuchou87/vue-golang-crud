package main

import (
	"github.com/yuchou87/vue-golang-crud/server/app"
	"github.com/yuchou87/vue-golang-crud/server/comm"
	"log"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS books
(
    id SERIAL PRIMARY KEY,
    title VARCHAR (50) NOT NULL,
    author VARCHAR (50) NOT NULL,
	status BOOLEAN DEFAULT false
)`

var a app.App

func main() {
	dbUser := comm.GetEnv("DB_USER", "postgres")
	dbPass := comm.GetEnv("DB_PASS", "test123456")
	dbName := comm.GetEnv("DB_NAME", "postgres")
	dbHost := comm.GetEnv("DB_HOST", "10.0.0.5")
	//dbReset := comm.GetEnv("DB_RESET", "false")
	a = app.App{}
	a.Initialize(dbUser, dbPass, dbName, dbHost)

	ensureTableExits()
	a.Run(":5000")
	//if dbReset == "true" {
	//	clearTable()
	//}
}

func ensureTableExits() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM books")
	a.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}
