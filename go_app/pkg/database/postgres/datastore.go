package postgres

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type db interface {
	Select(interface{}, string, ...interface{}) error
	Get(interface{}, string, ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	BeginTxx(context.Context, *sql.TxOptions) (*sqlx.Tx, error)
	Beginx() (*sqlx.Tx, error)
	Rebind(query string) string
	PingContext(context.Context) error
	Close() error
	MustBegin() *sqlx.Tx
}

type Dao struct {
	db
}

func NewDataStore(d db) *Dao {
	return &Dao{d}
}
