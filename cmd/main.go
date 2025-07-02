package main

import (
	"log"
	"net/http"

	"github.com/Dhanraj-Patil/post-comment-go/platform/router"
)

// @title CloudSEK - Post-Comment-Service
// @version 1.0
// @description Post-Comment-Service for CloudSEK Backend Intern Assignment.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	rtr := router.New()

	if err := http.ListenAndServe(":8080", rtr); err != nil {
		log.Fatalf("Failed to start server. Reason: %v", err)
	}
}
