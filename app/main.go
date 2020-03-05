package main

import (
  "log"
  "net/http"
)


func main() {

  http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  log.Println("Hello World!")
	}))

  log.Println("Webapp live on 3000.")
  http.ListenAndServe(":3000", nil)

}
