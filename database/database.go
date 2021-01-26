package database

// Database interface
type Database interface {
	Connection() error

	Ping() error

	GetColumnsInformation(TableName string) (map[string]interface{}, error)
}
