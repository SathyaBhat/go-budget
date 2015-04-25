package main

import (
	"database/sql"
	_ "fmt"
	"github.com/jcelliott/lumber"
	_ "github.com/mattn/go-sqlite3"
	_ "io/ioutil"
	"os"
)

var (
	log *lumber.ConsoleLogger
)

const DBNAME string = "db/budget.db"

func DbExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateTables() error {
	db, err := sql.Open("sqlite3", DBNAME)
	if err != nil {

		log.Fatal("Error: %s", err.Error())
	}

	createAccounts := `
create table accounts (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name string NOT NULL,
    start_bal INTEGER not NULL,
    start_date TIMESTAMP default CURRENT_TIMESTAMP,
    type INTEGER NOT NULL,
    created_by char(250),
    created_date TIMESTAMP not NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by char(250),
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(createAccounts, nil)
	defer db.Close()
	return err
}

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
