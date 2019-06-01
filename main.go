package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(rw, "Hello, world มิก้า!")
  })

  http.ListenAndServe(":8080", nil)
}