package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/brayanzuritadev/citas/models"
	_ "github.com/denisenkom/go-mssqldb"
)

var SQLDB *sql.DB

func ConnetionDb(ctx context.Context) error {

	user := ctx.Value(models.Key("username")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	port := 1433
	DatabaseName := ctx.Value(models.Key("database")).(string)
	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", host, user, password, port, DatabaseName)

	db, err := sql.Open("sqlserver", connStr)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Connection to SQL Server successful")
	SQLDB = db

	return nil
}

func CloseConnection() {
	if SQLDB != nil {
		fmt.Println("Closing connection DB")
		err := SQLDB.Close()
		if err != nil {
			fmt.Println("Closed connection DB")
			return
		}
	}
}

func PingConnection() bool {
	err := SQLDB.Ping()
	return err == nil
}
