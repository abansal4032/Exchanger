// Package dbclient provides interfaces and functionality to create
// database transactions and connections and run queries.
package dbclient

import (
	"database/sql"
	"database/sql/driver"
	"strings"

	_ "github.com/go-sql-driver/mysql" // loads the mysql driver.
)

// TransactionKey is the key used by context for transaction.
var TransactionKey Key = "transaction"

// Key defines the key type used by this package.
type Key string

// DB represents a database handle.
type DB interface {
	Begin() (Tx, error)
	Connection
}

// Tx represents a database transaction.
type Tx interface {
	driver.Tx
	Connection
}

// Connection represents a database connection.
type Connection interface {
	Execer
	Queryer
	Preparer
}

// Execer represents interface for executing queries.
type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// Preparer represents interface for preparing queries.
type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

// Queryer represents interface for queries to fetch data from data store.
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type myDBType struct {
	*sql.DB
}

func (db *myDBType) Begin() (Tx, error) {
	return db.DB.Begin()
}

// db is the database client.
var myDB DB

// Handle returns the database handle.
var Handle = func() DB {
	return myDB
}

var SqlDB *sql.DB

const DriverName = "mysql"

// NewTransaction starts and returns a new database transaction.
var NewTransaction = func() (Tx, error) {
	return myDB.Begin()
}

// MustNewTransaction starts and returns a new database transaction.
// It panics in case of an error.
var MustNewTransaction = func() Tx {
	tx, err := NewTransaction()
	if err != nil {
		panic(err)
	}
	return tx
}

// Connect establishes connection with the database.
var Connect = func(cfg DBConfig) error {
	tempDB, err := sql.Open(DriverName, dsn(cfg))
	if err != nil {
		return err
	}

	if err := tempDB.Ping(); err != nil {
		return err
	}

	tempDB.SetMaxIdleConns(cfg.MaxIdleConn)
	tempDB.SetMaxOpenConns(cfg.MaxOpenConn)

	SqlDB = tempDB
	myDB = &myDBType{tempDB}
	return nil
}

func dsn(cfg DBConfig) string {
	dataSourceName := cfg.User + ":" + cfg.Passwd + "@tcp(" + cfg.Host + ":3306)/" + cfg.Name + "?" + strings.Join(cfg.Params, "&")
	return dataSourceName
}

// Config defines the package configuration object.
type DBConfig struct {
	User        string   `json:"user"`
	Passwd      string   `json:"password"`
	Host        string   `json:"host"`
	Name        string   `json:"database"`
	Params      []string `json:"params"`
	MaxIdleConn int      `json:"maxIdleConns"`
	MaxOpenConn int      `json:"maxOpenConns"`
}
