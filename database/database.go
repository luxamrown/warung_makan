package database

import (
	"fmt"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

func Connect() {
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "warung_makan"
	dbUser := "postgres"
	dbPassword := "stauffenberg"

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	println(dataSourceName)
	conn, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		fmt.Println("Unable to connect to database: %v\n", err)
		panic(err)
	}
	defer conn.DB.Close()
}
