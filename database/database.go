package database

import (
	"authentication_microservice/model"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqlCreateUser = ` CREATE TABLE IF NOT EXISTS user(
        Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        Username VARCHAR NOT NULL
    );
	`
	sqlInsertUser = `
	INSERT INTO user 
		(Username) VALUES (?)
	`
	sqlFindUserByUsername = `
	SELECT * FROM user
		WHERE username = ?
	`
)

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("no database found")
	}
	err = migrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrate(db *sql.DB) error {

	_, err := db.Exec(sqlCreateUser)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(u *model.User, db *sql.DB) error {
	_, err := db.Exec(sqlInsertUser, u.Username)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func FindUserByUsername(username string, db *sql.DB) (*model.User, error) {
	getuser := model.User{}
	err := db.QueryRow(sqlFindUserByUsername, username).Scan(&getuser.Username)
	if err != nil {
		return nil, err
	}
	return &getuser, nil
}
