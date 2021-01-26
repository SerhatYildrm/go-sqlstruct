package commands

import (
	"go-sqlstruct/database"
	"os"

	"go-sqlstruct/templates"
)

// MSSQLExecute ...
func MSSQLExecute(tableName string, path string) error {
	var db database.Database
	var err error

	path = createDirectory(path)

	db = database.CreateMSSqlConnection("localhost", "store_sahin", "sa", "q1w2e3r4!", 0)
	if err = db.Connection(); err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	results, err := db.GetColumnsInformation(tableName)
	if err != nil {
		return err
	}

	var c *templates.SQLStruct
	c = templates.CreateSQLStruct(tableName, results)
	c.Create()
	c.WriteToGOFile(path)

	return err
}

func createDirectory(path string) string {
	path = path + "/Models"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0664)
	}
	return path
}
