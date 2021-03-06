package main

import (
	"encoding/csv"
	"flag"
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

var importPtr = flag.Bool("import", false, "Use -import to import csv")
var importFileName = flag.String("file", "", "Filename of the csv to be imported")

func fetchCSV(csvfilename string) (bool, error) {
	log.Debug("Import filename: %s",csvfilename )
	csvfile, err := os.Open(csvfilename)
    
	if err != nil {
		log.Error("Error %s occurred", err.Error())
		return false, err
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.Comma = ','
	reader.LazyQuotes = true
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
        log.Error("Error %s occurred", err.Error())
		return false, err
	}

	for _, each := range rawCSVdata {
		log.Info("Account: %s", each[0])
	}
	return true, err
}
func main() {


	log = lumber.NewConsoleLogger(lumber.DEBUG)
	log.Prefix("gbd")

	flag.Parse()

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

	if *importPtr {
		log.Debug("Importing CSV")
		log.Info("Not yet implemented")
		if *importFileName == "" {
			log.Error("Running in import mode, but filename not passed")
			panic("CSV file not found or not passed")
		}
		_, _ = fetchCSV(*importFileName)
	}

}
