package main

import (
    "fmt"
    "net/http"
    "os"
)


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html>")
    fmt.Fprintf(w, "<title>Welcome to DC/OS 101!</title><body>")
    fmt.Fprintf(w, "<h1>Welcome to DC/OS 101!</h1>")
    fmt.Fprintf(w, "<h1>Running on node '"+  os.Getenv("HOST") + "' and port '" + os.Getenv("PORT0") + "' </h1>")
}

func main() {
    fmt.Println("Starting DC/OS-101 App ")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":" + os.Getenv("PORT0") , nil)
}

