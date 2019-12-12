package main

import (
	"log"
	"os"
)

func (s *server) loadEnvVariables() {
	addr := os.Getenv("POSTGRES_ADDR")
	postgresPort := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")
	db := os.Getenv("POSTGRES_DB")
	port := os.Getenv("HTTP_PORT")

	if addr == "" {
		log.Fatal("POSTGRES_ADDR environment variable missing")
	}

	if postgresPort == "" {
		log.Fatal("POSTGRES_PORT environment variable missing")
	}

	if user == "" {
		log.Fatal("POSTGRES_USER environment variable missing")
	}

	if pass == "" {
		log.Fatal("POSTGRES_PASS environment variable missing")
	}

	if db == "" {
		log.Fatal("POSTGRES_DB environment variable missing")
	}

	if port == "" {
		log.Fatal("HTTP_PORT environment variable missing")
	}

	s.PostgresAddr = addr
	s.PostgresPort = postgresPort
	s.PostgresUser = user
	s.PostgresPass = pass
	s.PostgresDB = db
	s.Port = port
}
