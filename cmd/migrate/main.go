package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fahreyad/go_api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	conf := config.MustLoad()
	if len(os.Args) > 2 {
		log.Fatal("Too many arguments. Please provide only one argument: 'up' or 'down'.")
	}

	m, err := migrate.New(
		"file://migrations",
		conf.DATABASE_URL)

	if err != nil {
		log.Fatalf("migrate_new: Failed to create migrate instance: %v", err)
	}

	switch os.Args[1] {
	case "up":
		err = m.Up()
		if err != nil {
			log.Fatalf("migrate_up: Failed to run migrations up: %v", err)
		}
		fmt.Println("Running migrations up...")
	case "down":
		err = m.Down()
		if err != nil {
			log.Fatalf("migrate_down: Failed to run migrations down: %v", err)
		}
		fmt.Println("Running migrations down...")
		// Add your migration logic here for "down"
	default:
		fmt.Println("Invalid argument. Use 'up' or 'down'.")
	}

	fmt.Println("Running migrations...", os.Args[1])
}
