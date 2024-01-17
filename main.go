package main

import (
	"net/http"
)

func main() {
    setupJsonApi()

    server := &http.Server{
        Addr: ":80",
    }

    go func() {
        if err := server.ListenAndServe(); err != nil {
            fmt.Println("Server error:", err)
        }
    }()

    // Graceful shutdown on interrupt signal
    interruptSignal := make(chan os.Signal, 1)
    signal.Notify(interruptSignal, os.Interrupt)
    <-interruptSignal

    fmt.Println("Shutting down gracefully...")
    server.Shutdown(context.Background())
}
