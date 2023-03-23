package main

import (
	"fmt"

	"github.com/RianNegreiros/clean-go-rest-api/internal/comment"
	"github.com/RianNegreiros/clean-go-rest-api/internal/db"
	transportHttp "github.com/RianNegreiros/clean-go-rest-api/internal/transport/http"
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

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
