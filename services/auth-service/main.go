package main

import (
	"fmt"
	"hauparty/libs/common"
	"hauparty/libs/db"
	"log"
)

func main() {
	fmt.Println("Hello, Go!")
	common.LogInfo("Test")
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	fmt.Println(conn)
	// Use conn (gorm.DB instance)
	_ = conn
}
