package internal

import (
	"context"

	"github.com/amarantec/wallet/internal/db"
)

type Account struct {
  Id        uint      `json:"id"`
  Balance   float64   `json:"balance"`
}


func (a *Account) CreateAccount(ctx context.Context) (uint, error) {
  query := `
    INSERT INTO accounts (balance)
      VALUES (?);
  `
  stmt, err := db.DB.Prepare(query)
  if err != nil {
    return ZERO, err
  }
  defer stmt.Close()

  result, err := stmt.ExecContext(ctx, a.Balance)
  if err != nil {
    return ZERO, err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return ZERO, err
  }

  return uint(id), nil
}

func GetMyBalance(ctx context.Context, id uint) (float64, error){
  query := `
    SELECT balance FROM accounts WHERE id = ?;
  `

  row := db.DB.QueryRowContext(ctx, query, id)

  var balance float64
  err := row.Scan(&balance)
  if err != nil {
    return float64(ZERO), err
  }

  return balance, nil
}
