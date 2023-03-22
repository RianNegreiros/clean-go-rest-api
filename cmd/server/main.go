package main

import (
	"fmt"

	"github.com/RianNegreiros/clean-go-rest-api/internal/db"
)

// Run is the startup of the go application
func Run() error {
	fmt.Println("Starting the application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
