package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

type Account struct {
    Number string `json: "AccountNumber"`
    Balance string `json: "Balance"`
    Desc string `json: "AccountDescription"`
}

var Acounts []Account

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to our bank!")
  fmt.Println("Endpoint: /")
}
