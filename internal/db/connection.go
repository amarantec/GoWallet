package db

import (
  "database/sql"
  _"github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
  var err error
  DB, err = sql.Open("sqlite3", "api.db")

  if err != nil {
    panic("Could not connect to dabase.")
  }

  DB.SetMaxOpenConns(10)
  DB.SetMaxIdleConns(5)
  createTables()
}

func createTables() {
  createAccountTable := `
    CREATE TABLE IF NOT EXISTS accounts (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      balance REAL
  );
  `
  _, err := DB.Exec(createAccountTable)
  if err != nil {
    panic("Could not create accounts table.")
  }

  createUserTable := `
    CREATE TABLE IF NOT EXISTS users (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      email  TEXT NOT NULL UNIQUE,
      password TEXT NOT NULL
  );
  `
  _, err = DB.Exec(createUserTable)
  if err != nil {
    panic("Could not create user table.")
  }
}
