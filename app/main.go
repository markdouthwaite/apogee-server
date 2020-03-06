package main

import (
  "log"
  "net/http"
  "html/template"
  "os"
  "io/ioutil"
  "bytes"

  "github.com/gorilla/mux"
)

const HOST = ""

type Page struct {
  Title string
  Variables string
}


type QueryResponse struct {

}



func getVariables()(string){
  payload := []byte(`"evidence": {"PAP": "LOW"}`)
  response, err := http.Post("http://modelserver/inference/query", "application/json", bytes.NewBuffer(payload))

  if err != nil {
    panic(err)
  }

  defer response.Body.Close()

  body, err := ioutil.ReadAll(response.Body)

  println(body)

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
  println(getVariables())
  r := mux.NewRouter()
  r.HandleFunc("/app", HomeHandler)

  http.Handle("/", r)
  log.Println("Webapp live on 3000.")
  http.ListenAndServe(":3000", nil)

}
