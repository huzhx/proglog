package main

import (
	"github.com/huzhx/proglog/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8099")
	log.Fatal(srv.ListenAndServe())
}
