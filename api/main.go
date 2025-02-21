package handler

import (
  "fmt"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    fmt.Fprint("404 not found")
  }
  fmt.Fprintf(w, "Hello, from go on vercel!")
}
