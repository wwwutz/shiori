//go:generate go run assets-generator.go

package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/wwwutz/shiori/cmd"
	dt "github.com/wwwutz/shiori/database"
	fp "path/filepath"
)

var dataDir = "."

func main() {
	// Open database
	dbPath := fp.Join(dataDir, "shiori.db")
	sqliteDB, err := dt.OpenSQLiteDatabase(dbPath)
	checkError(err)

	// Start cmd
	shioriCmd := cmd.NewShioriCmd(sqliteDB, dataDir)
	if err := shioriCmd.Execute(); err != nil {
		logrus.Fatalln(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
