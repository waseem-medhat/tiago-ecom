package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wipdev-tech/tiago-ecom/cmd/api"
	"github.com/wipdev-tech/tiago-ecom/db"
)

var env = struct {
	dbURL   string
	dbToken string
}{}

func main() {
	godotenv.Load()
	env.dbURL = os.Getenv("DB_URL")
	env.dbToken = os.Getenv("DB_TOKEN")

	db, err := db.NewDB(env.dbURL, env.dbToken)
	if err != nil {
		log.Fatal("couldn't open DB connection:", err)
	}
	defer db.Close()

	s := api.NewAPIServer(":8080", db)
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
