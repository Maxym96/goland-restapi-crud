package postgresql

import (
	"github.com/joho/godotenv"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"os"
)

func NewClient() (db *db.Session, err error) {

	if err = godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	BdHost := os.Getenv("BD_HOST")
	BdDatabase := os.Getenv("BD_DATABASE")
	BdUser := os.Getenv("BD_USER")
	BdPassword := os.Getenv("BD_PASSWORD")

	settings := postgresql.ConnectionURL{
		Database: BdDatabase,
		Host:     BdHost,
		User:     BdUser,
		Password: BdPassword,
	}
	sess, err := postgresql.Open(settings)

	if err != nil {
		log.Fatal("postgresql.Open: ", err)
	}
	//defer sess.Close()

	if err := sess.Ping(); err != nil {
		log.Fatal("Ping: ", err)
	}
	log.Printf("Successfully connected to database: %q", sess.Name())

	return &sess, nil
}
