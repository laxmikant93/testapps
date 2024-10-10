package main

import (
    "fmt"
    "log"
    "net/http"
    "goauth/handler"
    "github.com/rs/cors"
)

func main() {
    // Create a new CORS handler with your desired configuration
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"}, // Allow all origins, or specify a list
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })

    // Routes
    http.HandleFunc("/signup", handler.SignUp)
    http.HandleFunc("/signin", handler.SignIn)

    // Wrap the router with the CORS middleware
    handler := c.Handler(http.DefaultServeMux)

    // Start server
    fmt.Println("Server started on :8000")
    log.Fatal(http.ListenAndServe(":8000", handler))
}
