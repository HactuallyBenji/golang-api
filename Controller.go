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

var Accounts []Account

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to our bank!")
  fmt.Println("Endpoint: /")
}

func returnAllAccounts(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(Accounts)
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/accounts", returnAllAccounts)
  log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
  Accounts = []Account{ 
    Account{Number: "C45t34534", Balance: "24545.5", Desc: "Checking Account"},
		Account{Number: "S3r53455345", Balance: "444.4", Desc: "Saving Account"},
	}

  handleRequests()
}
