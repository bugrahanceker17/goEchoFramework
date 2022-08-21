package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	username := os.Getenv("USER")
	pasword := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(pasword)
}
