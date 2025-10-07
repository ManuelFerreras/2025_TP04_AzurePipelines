// back/main.go
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
  })

  port := os.Getenv("PORT")
  if port == "" { port = "8080" }

  log.Printf("listening on :%s", port)
  log.Fatal(http.ListenAndServe(":"+port, mux))
}
