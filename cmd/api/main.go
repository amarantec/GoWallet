package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/amarantec/wallet/internal/db"
	"github.com/amarantec/wallet/internal/handlers"
)

func main() {
  db.InitDB()

  handlers.LoadTemplates()
  mux := handlers.SetRoutes()

  mux.Handle("/css/",
    http.StripPrefix("/css/",
      http.FileServer(http.Dir("../../web/css"))))

  server := &http.Server{
    Addr: "127.0.0.1:8080",
    Handler: mux,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
  }

  fmt.Printf("Server listen on: %s\n", server.Addr)
  log.Fatal(server.ListenAndServe())
}
