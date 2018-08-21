package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from golang http server")
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe("0.0.0.0:9898", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
