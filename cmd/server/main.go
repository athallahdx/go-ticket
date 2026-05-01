package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"go-ticket/internal/config"
	"go-ticket/internal/handler"
	"go-ticket/internal/repository"
	"go-ticket/internal/service"

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

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)

	userHandler := handler.NewUserHandler(userService, cfg)
	authHandler := handler.NewAuthHandler(authService, cfg)

	fmt.Println("✅ MySQL connected Successfully!")
	fmt.Printf("✅ Server running on Port %s...\n", cfg.Port)

	router := SetupRouter(
		userHandler,
		authHandler,
		cfg,
	)

	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
