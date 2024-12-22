package internal

import (
	"context"
	"fmt"

	"github.com/amarantec/wallet/internal/db"
	"github.com/amarantec/wallet/internal/utils"
)

type User struct {
  Id        uint    `json:"id"`
  Email     string  `json:"email"`
  Password  string  `json:"password"`
}

func (u *User) Register(ctx context.Context) (uint, error) {
  query := `
    INSERT INTO users (email, password) VALUES (?, ?);
  `
  stmt, err := db.DB.Prepare(query)
  if err != nil {
    return ZERO, err
  }

  defer stmt.Close()
 
  hashedPassword, err := utils.HashPassword(u.Password)
  if err != nil {
    return ZERO, err
  }

  result, err := stmt.ExecContext(ctx, u.Email, hashedPassword)
  if err != nil {
    return ZERO, err
  }

  userId, _ := result.LastInsertId()

  return uint(userId), nil
}

func (u *User) Login(ctx context.Context) (uint, error) {
  var retriviedPassword string

  query := `
    SELECT id, password FROM users WHERE email = ?;
  `
  row := db.DB.QueryRowContext(ctx, query, u.Email)

  err := row.Scan(&u.Id, &retriviedPassword)
  if err != nil {
    return ZERO, err
  }

  passwordIsValid := utils.CheckPassword(u.Password, retriviedPassword)
  if !passwordIsValid {
    return ZERO, fmt.Errorf("invalid password")
  }

  return u.Id, nil
}
