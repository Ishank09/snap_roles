package db

import (
	"context"
	"database/sql"
	"example/snap_roles/cmd/model"
	"example/snap_roles/configs"
	"fmt"
	"log"
)

type AzureSqlInterface interface {
	DatabaseConnection()
}

type AzureSqlStruct struct {
	Config configs.ConfigInterface
}

func getDbConnection(appConfig configs.AppConfig) model.Db_Connection {
	var db_connection model.Db_Connection
	db_connection.DatabaseName = appConfig.DB_NAME
	db_connection.UserId = appConfig.DB_USER
	db_connection.Server = appConfig.DB_SERVER
	db_connection.Port = appConfig.DB_PORT
	db_connection.Password = appConfig.DB_PASSWORD
	return db_connection
}

func (a *AzureSqlStruct) DatabaseConnection() {
	var db_connection_model model.Db_Connection
	var db *sql.DB
	appconfig, err01 := a.Config.LoadConfig()
	if err01 != nil {
	}
	db_connection_model = getDbConnection(*appconfig)

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		db_connection_model.Server, db_connection_model.UserId, db_connection_model.Password, db_connection_model.Port, db_connection_model.DatabaseName)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
}
