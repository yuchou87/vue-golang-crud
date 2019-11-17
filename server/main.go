package main

import (
	"github.com/yuchou87/vue-golang-crud/server/app"
	"github.com/yuchou87/vue-golang-crud/server/comm"
)

func main() {
	dbUser := comm.GetEnv("DB_USER", "postgres")
	dbPass := comm.GetEnv("DB_PASS", "test123456")
	dbName := comm.GetEnv("DB_NAME", "postgres")
	dbHost := comm.GetEnv("DB_HOST", "10.0.0.5")
	a := app.App{}
	a.Initialize(dbUser, dbPass, dbName, dbHost)
	a.Run(":5000")
}
