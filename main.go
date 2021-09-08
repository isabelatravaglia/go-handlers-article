package main

import (
  "fmt"
  "net/http"
)

type HandlePage struct {}

func (h HandlePage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Welcome to the handle page!</h1>")
}

func HandleFuncPage(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Welcome to the handlefunc page!</h1>")
}

func main() {
  NewHandlePage := new(HandlePage)

  http.Handle("/handlepage", NewHandlePage)
  http.HandleFunc("/handlefuncpage", HandleFuncPage)
  http.ListenAndServe(":3000", nil)
}
