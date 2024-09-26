package repository

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/purisaurabh/database-connection/config"
)

type Repository interface {
}

type RepositoryStruct struct {
	DB *sql.DB
}

func Init() (RepositoryStruct, error) {
	dbConfig := config.Database()
	db, err := sql.Open(dbConfig.Driver(), dbConfig.ConnectionURL())
	if err != nil {
		fmt.Println("Error in opening connection : ", err)
		return RepositoryStruct{}, err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Error in pinging the connection : ", err)
		return RepositoryStruct{}, err
	}

	db.SetMaxIdleConns(dbConfig.MaxPoolSize())
	db.SetMaxOpenConns(dbConfig.MaxOpenConns())
	db.SetConnMaxLifetime(time.Duration(dbConfig.MaxLifeTimeMins()) * time.Minute)

	return RepositoryStruct{DB: db}, nil
}
