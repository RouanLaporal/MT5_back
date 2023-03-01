package main

import (
	database "back_project/mysql"
	"back_project/web"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func enableCors(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")

}

func main() {
	conf := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "database:3306",
		DBName:               "MT5_back_project",
		AllowNativePasswords: true, // Il faut le préciser
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	store := database.CreateStore(db)
	mux := web.NewHandler(store)

	err = http.ListenAndServe(":8097", mux)
	if err != nil {
		_ = fmt.Errorf("impossible de lancer le serveur : %w", err)
		return
	}

}
