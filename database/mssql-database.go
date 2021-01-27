package database

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type MSSql struct {
	host   string
	port   int
	dbname string

	user     string
	password string

	dbconn *sql.DB
}

// CreateMSSqlConnection ...
func CreateMSSqlConnection(host, dbname, user, password string, port int) *MSSql {
	return &MSSql{host, port, dbname, user, password, nil}
}

// Generate connection string
func (m *MSSql) connectionString() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", m.user, m.password, m.host, m.dbname)
}

// Connection ...
func (m *MSSql) Connection() error {
	var err error
	m.dbconn, err = sql.Open("sqlserver", m.connectionString())
	if err != nil {
		return err
	}
	return nil
}

// Ping ...
func (m *MSSql) Ping() error {
	if err := m.dbconn.Ping(); err != nil {
		return err
	}
	return nil
}

// GetColumnsInformation ...
// https://stackoverflow.com/a/29360030/11565617
func (m *MSSql) GetColumnsInformation(TableName string) (map[string]interface{}, error) {
	var results = make(map[string]interface{})

	rows, err := m.dbconn.Query("SELECT TOP 1 * FROM " + TableName)
	defer rows.Close()

	colNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	cols := make([]interface{}, len(colNames))
	colPtr := make([]interface{}, len(colNames))

	for i := 0; i < len(cols); i++ {
		colPtr[i] = &cols[i]
	}

	for rows.Next() {
		err = rows.Scan(colPtr...)
		if err != nil {
			return nil, err
		}

		colTypes, err := rows.ColumnTypes()
		if err != nil {
			return nil, err
		}
		/*for i, col := range cols {
			//results[colNames[i]] = reflect.TypeOf(col)
			results[colNames[i]] = colTypes[i].DatabaseTypeName()
		}*/

		var sqlType string
		for i := 0; i < len(cols); i++ {
			sqlType = colTypes[i].DatabaseTypeName()
			results[colNames[i]] = ConvertColumnType(sqlType)
		}
	}

	return results, nil
}

// ConvertColumnType ...
func ConvertColumnType(columnType string) string {
	if columnType == "INT" || columnType == "BIT" {
		return "int"
	} else if columnType == "NVARCHAR" {
		return "string"
	} else if columnType == "DATETIME" {
		return "time.Time"
	}
	return "interface{}"
}
