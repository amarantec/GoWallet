package utils

import "net/http"

func HandleError(w http.ResponseWriter, msg string, err error, status int) bool {
  if err != nil {
    http.Error(w,
      msg + ". error: " + err.Error(),
      status,
      )
    return true
  }
  return false
}
