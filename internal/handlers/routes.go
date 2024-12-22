package handlers

import "net/http"

func SetRoutes() *http.ServeMux{
  m := http.NewServeMux()

  m.HandleFunc("/home", index) 
  m.HandleFunc("/api/register", register)
  m.HandleFunc("/api/login", login)
  m.HandleFunc("/api/create-account", createAccount)
  m.HandleFunc("/api/get-my-balance/{id}", getMyBalance)

  return m
}
