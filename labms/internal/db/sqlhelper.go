package db

import (
	"repogin/internal/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

type SQLRepo struct {
	Session *dbr.Session
	DBInfo  models.DBInfo
}

func NewSQLRepo(sm models.DBInfo) *SQLRepo {

	DB, DError := dbr.Open("mysql", sm.DSN, nil)
	if DError != nil {
		fmt.Println("ERROR : NewSQLConnection")
	}
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(5)
	return &SQLRepo{
		Session: DB.NewSession(nil),
		DBInfo:  sm,
	}
}
