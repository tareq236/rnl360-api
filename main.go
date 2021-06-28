package main

import (
	"fmt"
	"net/http"
	"os"
	"rnl360-api/routes"

	"github.com/joho/godotenv"
)

func main() {

	// .env file load
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	// get port information from the env file
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("env error:", "PORT not initialize in the .env file.")
	}

	fmt.Println("server started and listening on port: 0.0.0.0:", port)

	router := routes.SetupRouter()

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}

}
