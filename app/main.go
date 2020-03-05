package main

import (
  "log"
  "net/http"
  "html/template"
  "os"
  // "bytes"

  "github.com/gorilla/mux"
)

const HOST = ""

type Page struct {
  Title string
  Variables string
}



func getVariables() (string){
  // payload := []byte(``)
  // resp, _ := http.Post("http://localhost:4504/inference/query", "application/json", bytes.NewBuffer(payload))

  return "test"
}



func HomeHandler(w http.ResponseWriter, r *http.Request) {

  w.WriteHeader(http.StatusOK)

  wd, _ := os.Getwd()

  page := Page{Title: "Home", Variables: getVariables()}
  path := wd + "/templates/index.html"

  t, _ := template.ParseFiles(path)
  t.Execute(w, page)
}



func main() {
  r := mux.NewRouter()
  r.HandleFunc("/app", HomeHandler)

  http.Handle("/", r)
  log.Println("Webapp live on 3000.")
  http.ListenAndServe(":3000", nil)

}
