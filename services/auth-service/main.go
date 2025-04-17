package main

import (
	"fmt"
	"hauparty/libs/common/utils"
	"hauparty/libs/db"
	"log"
)

func main() {
	fmt.Println("Hello, Go!")
	utils.LogInfo("Test")
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	fmt.Println("Connection", conn)
	// Use conn (gorm.DB instance)
	_ = conn
	InitApp()
}
