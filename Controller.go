package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "io/ioutil"
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

func returnAccount(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Received returnAccount request")
  fmt.Println(r.Method)
  vars := mux.Vars(r)
  key := vars["number"]
  for _, account := range Accounts {
    if account.Number == key {
      json.NewEncoder(w).Encode(account)
    }
  }
}

func createAccount(w http.ResponseWriter, r *http.Request) {
  reqBody, _ := ioutil.ReadAll(r.Body)
  var account Account
  json.Unmarshal(reqBody, &account)
  Accounts = append(Accounts, account)
  json.NewEncoder(w).Encode(account)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
    // use mux to parse the path parameters
    vars := mux.Vars(r)
    // extract the account number of the account we wish to delete
    id := vars["number"]
    // we then need to loop through dataset
    for index, account := range Accounts {
        // if our id path parameter matches one of our
        // account numbers
        if account.Number == id {
            // updates our dataset to remove the account
            Accounts = append(Accounts[:index], Accounts[index + 1:]...)
        }
    }
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["number"]

  reqBody, _ := ioutil.ReadAll(r.Body)
  var newAccount Account
  json.Unmarshal(reqBody, &newAccount)

  for index, account := range Accounts {
    if account.Number == id {
      Accounts[index].Number = newAccount.Number
      Accounts[index].Balance = newAccount.Balance
      Accounts[index].Desc = newAccount.Desc
    }
  }
}

func handleRequests() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", homePage)
  router.HandleFunc("/accounts", returnAllAccounts)
  router.HandleFunc("/account/{number}", returnAccount).Methods("GET")
  router.HandleFunc("/account", createAccount).Methods("POST")
  router.HandleFunc("/account/{number}", deleteAccount).Methods("DELETE")
  router.HandleFunc("/account/{number}", updateAccount).Methods("PUT")
  log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
  Accounts = []Account{ 
    Account{Number: "C45t34534", Balance: "24545.5", Desc: "Checking Account"},
		Account{Number: "S3r53455345", Balance: "444.4", Desc: "Saving Account"},
	}

  handleRequests()
}
