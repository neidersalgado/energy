package db

import (
	"database/sql"
	"fmt"
	"github.com/energy/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewDBConn() (*sql.DB, error) {
	conf := config.GetConfig()
	
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)
	
	db, err := sql.Open("mysql", dataSourceName)
	
	if err != nil {
		return nil, err
	}
	
	return db, nil
}
