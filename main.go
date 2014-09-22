package main

import (
  "fmt"
  "net/http"
  "strings"
)


var data = make(map[string]string)


func setHandler(w http.ResponseWriter, r *http.Request) {
  key := strings.Split(r.URL.Path[1:], "/")[1]
  value := strings.Split(r.URL.Path[1:], "/")[2]
  data[key] = value
}

func getHandler(w http.ResponseWriter, r *http.Request) {
  key := strings.Split(r.URL.Path[1:], "/")[1]
  fmt.Fprintf(w, data[key])
}

func main() {
  
  data["hugo"] = "hello!"
  
  
  http.HandleFunc("/set/", setHandler)
  http.HandleFunc("/get/", getHandler)
  http.ListenAndServe(":9090", nil)
}