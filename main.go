package main

import (
  "fmt"
  "net/http"
  "log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
  fmt.Println("hello from golang http server")
}

func main() {
  http.HandleFunc("/", sayHelloName)
  err := http.ListenAndServe("0.0.0.0:9898", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
