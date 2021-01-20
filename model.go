// model.go

package main

import (
    "fmt"
    "database/sql"
    "errors"
)

type user struct {
    Id    int     `json:"id"`
    Name  string  `json:"name"`
    Email string  `json:"email"`
}

func (u *user) getUser(db *sql.DB) error {
  return db.QueryRow("SELECT name, email FROM user WHERE id=$1",
        u.Id).Scan(&u.Name, &u.Email)
  //return errors.New("Not implemented")
}

func (u *user) updateUser(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (u *user) deleteUser(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (u *user) createUser(db *sql.DB) error {
  /*err := db.QueryRow(
    "INSERT INTO user(name, email) VALUES($1, $2) RETURNING id",
    u.Name, u.Email).Scan(&u.Id) */
    query := "INSERT INTO `User`(name, email) VALUES ('%s', '%s');"
    fmt.Printf("Query : %s", fmt.Sprintf(query, u.Name, u.Email))
    _ , err := db.Query(fmt.Sprintf(query, u.Name, u.Email))

  if err != nil {
      return err
  }
  return nil
//return errors.New("Not implemented")
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
  /*rows, err := db.Query(
    "SELECT id, name,  email FROM user LIMIT $1 OFFSET $2", count, start)*/
  
  rows, err := db.Query("SELECT * FROM user")
  
  if err != nil {
      return nil, err
  }

  defer rows.Close()
  users := []user{}

  for rows.Next() {
      var u user
      if err := rows.Scan(&u.Id, &u.Name, &u.Email); err != nil {
          return nil, err
      }
      users = append(users, u)
    }

  return users, nil
//return nil, errors.New("Not implemented")
}
