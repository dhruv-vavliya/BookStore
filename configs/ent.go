package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dhruv-vavliya/BookStore/ent"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	client *ent.Client
)

func Connect() (*ent.Client, error) {

	err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println(dbURI)

	client, err := ent.Open(os.Getenv("DATABASE"), dbURI, ent.Debug(), ent.Log(func(i ...interface{}) {
		for _,v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"),v)
			fmt.Print("\n")
		}
	}))

	if err != nil {
		log.Fatalf("connection failed", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err
}

