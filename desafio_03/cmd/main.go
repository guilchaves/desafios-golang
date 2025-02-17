package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/guilchaves/desafios-golang/desafio_03/configs"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const seedFilePath = "./import.sql"

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = seedDatabase(db)
	if err != nil {
		panic(fmt.Sprintf("failed to seed database: %v", err))
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":8080", nil)

}

func seedDatabase(db *gorm.DB) error {
	dropTables := `
		DROP TABLE IF EXISTS clients;
	`
	err := db.Exec(dropTables).Error
	if err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	log.Println("Tables dropped successfully")
	log.Println("Running migrations...")

	db.AutoMigrate(&entity.Client{})
	log.Println("Migrations completed.")

	log.Println("Seeding db from import.sql")
	fileContent, err := os.ReadFile(seedFilePath)
	if err != nil {
		return fmt.Errorf("failed to read import.sql %w", err)
	}

	statements := strings.Split(string(fileContent), ";")
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		log.Printf("Executing: %s", stmt)
		err := db.Exec(stmt).Error
		if err != nil {
			log.Printf("Failed to execute: %s\nError: %v", stmt, err)
		}
	}

	log.Println("Database seeding completed")
	return nil
}
