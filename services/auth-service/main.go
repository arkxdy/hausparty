package main

import (
	"context"
	"fmt"
	"hausparty/libs/db"
	"log"
	"os"
)

func main() {

	ctx := context.Background()
	println("service name", os.Getenv("SERVICE_NAME"))
	dbFactory, err := db.Connect(ctx)

	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	// For identity service:
	if os.Getenv("SERVICE_NAME") == "identity" {
		err = dbFactory.AutoMigrateIdentity()
	}
	if err != nil {
		log.Fatalf("Auto-migrate error: %v", err)
	}

	db := dbFactory.GetPostgres()
	fmt.Println("Connection", db)
	InitApp()
}
