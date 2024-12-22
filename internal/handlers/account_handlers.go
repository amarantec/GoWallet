package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/wallet/internal"
)

func createAccount(w http.ResponseWriter, r *http.Request) {
  account := internal.Account{}
  if err := 
    json.NewDecoder(r.Body).Decode(&account); err != nil {
      http.Error(w,
      "could not decote this request. error: " + err.Error(),
        http.StatusBadRequest,
      )
    return
  }

  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  response, err := account.CreateAccount(ctxTimeout)
  if err != nil {
    http.Error(w,
      "could not create this account. error: " + err.Error(),
      http.StatusInternalServerError,
      )
    return
  }
  
  jsonResponse, _ := json.Marshal(response)

  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonResponse)
}

func getMyBalance(w http.ResponseWriter, r *http.Request) {
  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  id, err := strconv.Atoi(r.PathValue("id"))
  if err != nil {
    http.Error(w,
      "invalid parameter. error: " + err.Error(),
      http.StatusBadRequest)
    return
  }

  balance, err := internal.GetMyBalance(ctxTimeout, uint(id))
  if err != nil {
    http.Error(w,
      "could not get the balance. error: " + err.Error(),
      http.StatusInternalServerError,
      )
    return
  }

  jsonResponse, _ := json.Marshal(balance)

  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonResponse)
}
