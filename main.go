package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var RunMode string
var IsLocal bool

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	RunMode = os.Getenv("RUN_MODE")
	IsLocal = RunMode == "dev"
}
func main() {
	fmt.Println(RunMode)
	fmt.Println(IsLocal)
	fmt.Println("hello word!")
}
