package bun

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DBClient struct {
	DB *bun.DB
}

func InitDatabase(server, postgres, opts string) (*DBClient, error) {
	connString := fmt.Sprintf("%s%s", server, opts)
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connString)))
	db := bun.NewDB(sqlDB, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true), // log everything
	))
	//verify connection
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	//check for database
	query := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname='%s';", postgres)
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		createDBquery := fmt.Sprintf("CREATE DATABASE %s;", postgres)

		_, err = db.Exec(createDBquery)
		if err != nil {
			return nil, err
		}
	}
	connString = fmt.Sprintf("%s%s", server, opts)
	sqlDB = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connString)))
	db = bun.NewDB(sqlDB, pgdialect.New())

	//verify connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DBClient{
		DB: db,
	}, nil

}

func (db *DBClient) Close() error {
	return db.Close()
}
