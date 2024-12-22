package handlers

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
  err := Templates.ExecuteTemplate(w, "index.html", nil)
  if err != nil {
    http.Error(w, "could no execute template: " + err.Error(), http.StatusInternalServerError)
    return
  }
}
