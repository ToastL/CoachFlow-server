package main

import (
    "log"
    "coachflow/internal/app"
)

func main() {
    server := app.NewApp()
    if err := app.Run(server); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}