package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

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
	if err == nil {
		log.Debug("Accounts table created")
	} else {
		log.Fatal("Error %s while creating accounts table", err.Error())
	}

	createPayees := `
create table payees (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name string NOT NULL,
    location string,
    created_by char(250),
    created_date TIMESTAMP not NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by char(250),
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(createPayees, nil)
	if err == nil {
		log.Debug("Payees table created")
	} else {
		log.Fatal("Error %s while creating Payees table", err.Error())
	}
	createTransactions := `
create table transactions (
    id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
    date timestamp NOT NULL,
    payee_id integer NOT NULL,
    category_id integer NOT NULL,
    details string(1000) NOT NULL,
    amount int NOT NULL,
    cleared_flag integer,
    created_by char(250),
    created_date timestamp NOT NULL DEFAULT current_timestamp,
    updated_by char(250),
    updated_date timestamp DEFAULT current_timestamp
    );`
	_, err = db.Exec(createTransactions, nil)
	if err == nil {
		log.Debug("Transactions table created")
	} else {
		log.Fatal("Error %s while creating Transactions table", err.Error())
	}
	defer db.Close()

	return err
}
