package postgres

import "github.com/jmoiron/sqlx"

// Postgres ...
type Postgres struct {
	db *sqlx.DB
}

// Initdb ...
func Initdb(psqlConfig string) (*Postgres, error) {
	var err error

	tempDB, err := sqlx.Connect("postgres", psqlConfig)

	if err != nil {
		return nil, err
	}

	return &Postgres{
		db: tempDB,
	}, nil
}
