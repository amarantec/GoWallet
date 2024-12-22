package handlers

import (
	"log"
	"text/template"
)

var Templates *template.Template

func LoadTemplates() {
  var err error
  Templates, err = template.ParseGlob("../../web/templates/*.html")
  if err != nil {
    log.Fatalf("Error loading templates: %v\n", err)
  }
}
