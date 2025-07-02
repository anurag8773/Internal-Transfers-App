package main

import (
	"log"
	"net/http"

	"internal-transfers/database"
	"internal-transfers/router"
)

func main() {
	db.InitDB()
	r := router.SetupRouter()

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
