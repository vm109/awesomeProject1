package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostgresUser string

const (
	postgresConnectionStringPattern = "postgres://%s:%s@%s:%s/%s"
	RootPostgresUser                = "postgres"
	RootPassword                    = "BH09gf2430%23%21"
	PORT                            = "5432"
	RootDatabase                    = "postgres"
	HOST                            = "db.reqpokkztpjowzlugnng.supabase.co"
)

func postgresConnectionString() (value string) {

	value = fmt.Sprintf(postgresConnectionStringPattern,
		RootPostgresUser,
		RootPassword,
		HOST,
		PORT,
		RootDatabase,
	)

	fmt.Println(value)
	return
}
func CreateSupaClient() *Dao {
	return NewDataStore(sqlx.MustOpen("postgres", postgresConnectionString()))
}
