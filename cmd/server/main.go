package main

import (
    "log"
    "github.com/asitm9/proglog/internal/server"
)

func main() {
    srv := server.NewHTTPServer(":8999")
    log.Fatal(srv.ListenAndServe())
}

