package main

import (
	"database/sql"
	"fmt"
	"log"

	"go-ticket/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error open database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database can't be reached: %v", err)
	}

	fmt.Println("✅ MySQL connected Successfully!")
	fmt.Printf("✅ Server running on Port %s...\n", cfg.Port)
}
