package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

func (s *server) connectToDB() {
	conn := "postgres://" + s.PostgresUser + ":" + s.PostgresPass + "@" + s.PostgresAddr + ":" + s.PostgresPort + "/" + s.PostgresDB + "?sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatalf("Malformed PostgreSQL connection string, details: %v", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Cannot connect to PostgreSQL instance, details: %v", err)
	}

	s.DB = db
}

func (s *server) checkSchema() {
	res, err := s.DB.Query("SELECT to_regclass('users')")
	defer res.Close()

	if err != nil {
		log.Fatalf("Error checking for schema, details: %v", err)
	}

	if res.Next() {
		var row string

		res.Scan(&row)

		if row == "" {
			log.Println("Users table doesn't exist, creating...")

			_, err2 := s.DB.Exec("CREATE TABLE users(name VARCHAR(50) PRIMARY KEY NOT NULL, color VARCHAR(30) NOT NULL, cats BOOLEAN NOT NULL)")

			if err2 != nil {
				log.Fatalf("Cannot create users table, details: %v", err2)
			}

			log.Println("Users table created successfully...")
		} else {
			log.Println("Users table already exists, skipping creation...")
		}
	}
}

func (s *server) SaveData(name, color string, cats bool) (bool, error) {
	res, err := s.DB.Query(`SELECT COUNT(name) FROM users WHERE name=$1`, name)

	if err != nil {
		log.Fatalf("Couldn't query database, details: %v", err)
	}

	for res.Next() {
		var count int

		if err = res.Scan(&count); err != nil {
			log.Fatal(err)
		}

		if count != 0 {
			return false, errors.New("A user with that name already exists in DB")
		}

		createRes, err1 := s.DB.Exec(`INSERT INTO users VALUES($1,$2,$3)`, name, color, cats)

		if err1 != nil {
			return false, err1
		}

		total, err2 := createRes.RowsAffected()

		if err2 != nil {
			return false, err2
		}

		if total == 1 {
			return true, nil
		}
	}

	return false, errors.New("Unexpected error")
}
