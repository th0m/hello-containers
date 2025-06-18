package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

    message := os.Getenv("MESSAGE")
    instanceId := os.Getenv("CLOUDFLARE_DEPLOYMENT_ID")

    log.Printf("Environment MESSAGE: %s", message)
    log.Printf("Environment CLOUDFLARE_DEPLOYMENT_ID: %s", instanceId)

    fmt.Fprintf(w, "Hi, I'm a container and this is my message: %s, and my instance ID is: %s", message, instanceId)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
    panic("This is a panic")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/container", handler)
    http.HandleFunc("/error", errorHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
