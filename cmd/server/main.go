package main

import "fmt"

// Run is the startup of the go application
func Run() error {
	fmt.Println("Starting the application")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
