package main

import (
	_ "fmt"
	"github.com/jcelliott/lumber"
	_ "github.com/mattn/go-sqlite3"
	_ "io/ioutil"
	_ "os"
)

var (
	log *lumber.ConsoleLogger
)

const DBNAME string = "db/budget.db"


func main() {
	log = lumber.NewConsoleLogger(lumber.DEBUG)
	log.Prefix("gbd")
	createdb, _ := DbExists(DBNAME)

	if !createdb {
		log.Info("DB not found, creating")
		err := CreateTables()
		if err != nil {
			log.Error("Error occured while creating tables: %s", err.Error())
			panic("unexpected error " + err.Error())
		}
	} else {
		log.Info("DB found, not creating")

	}

}
