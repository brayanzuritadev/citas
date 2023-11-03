package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/brayanzuritadev/citas/models"
	_ "github.com/denisenkom/go-mssqldb"
)

var SQLDB *sql.DB
var DatabaseName string

func ConnetionDb(ctx context.Context) error {

	user := ctx.Value(models.Key("username")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	port := 1433
	DatabaseName = ctx.Value(models.Key("database")).(string)
	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", host, user, password, port, DatabaseName)

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Connection to SQL Server successful")
	SQLDB = db
	return nil
}

func CloseConnection() {
	if SQLDB != nil {
		SQLDB.Close()
	}
}

func PingConnection() bool {
	err := SQLDB.Ping()
	return err == nil
}
