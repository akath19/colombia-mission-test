package main

import (
	"database/sql"
	"log"
	"net/http"
)

type server struct {
	PostgresAddr string
	PostgresPort string
	PostgresUser string
	PostgresPass string
	PostgresDB   string
	DB           *sql.DB
	Port         string
}

func main() {
	log.Println("Colombia Mission Technical Test Starting...")

	s := server{}

	log.Println("Loading & validating environment variables...")

	s.loadEnvVariables()

	log.Printf("Connecting to DB located at %v", s.PostgresAddr)

	s.connectToDB()

	defer s.DB.Close()

	s.checkSchema()

	log.Println("Schema verification finished, starting HTTP server...")

	http.HandleFunc("/", s.showForm)
	log.Fatal(http.ListenAndServe(":"+s.Port, nil))
}
