package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/amarantec/wallet/internal"
	"github.com/amarantec/wallet/internal/utils"
)

func register(w http.ResponseWriter, r *http.Request) {
  user := internal.User{}
  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
    http.Error(w,
      "could not decode this request. error: " + err.Error(),
        http.StatusBadRequest,
    )
    return
  }

  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  response, err := user.Register(ctxTimeout)
  if utils.HandleError(w, "could not register", err, http.StatusInternalServerError) {
    return
  }

  jsonResponse, _ := json.Marshal(response)

  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  w.WriteHeader(http.StatusCreated)
  w.Write(jsonResponse)

}

func login(w http.ResponseWriter, r *http.Request) {
  user := internal.User{}
  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
    http.Error(w,
      "could not decode this request. error: " + err.Error(),
      http.StatusBadRequest,
      )
    return
  }

  ctxTimeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  response, err := user.Login(ctxTimeout)
  if utils.HandleError(w, "could not login", err, http.StatusInternalServerError) {
    return
  }

  jsonResponse, _ := json.Marshal(response)

  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonResponse)
}
