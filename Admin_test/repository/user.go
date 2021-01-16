package repository

import (
	model "awesomeProject3/model"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)
import _ "github.com/go-sql-driver/mysql"

type Repository struct {
	Conn *sqlx.DB
}

func (repo Repository) CheckToken(token string) (model.Authorization, error) {
	var tokenTest model.TokenTest
	err := repo.Conn.Get(&tokenTest, "select user_id, token, expire_time from session where expire_time >= sysdate() and token = ?", token)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Wrap(err, "error connect DB")
	}
	if err != nil && err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}

func (repo Repository) GetUser() ([]model.User, error) {
	var userList []model.User
	err := repo.Conn.Select(&userList, "select iduser, login, password from user")
	if err != nil {
		return nil, errors.Wrap(err, "error select user")
	}

	return userList, nil
}

func (repo Repository) AddUser(login string, password string) error {
	_, err := repo.Conn.Exec("insert into user (login, password) values(? , ?)", login, password)

	if err != nil {
		return errors.Wrap(err, "error add user")
	}
	return nil
}
func (repo Repository) GetOwnUser(login string, password string, token string) ([]model.TokenTest, error) {
	var tokenTest []model.TokenTest
	var userList []model.User
	err := repo.Conn.Select(&userList, "select iduser, login, password from user where login = ? and password = ?", login, password)
	if err != nil {
		return nil, errors.Wrap(err, "error login user")
	} else {
		_, err := repo.Conn.Exec("insert into session (user_id, token, expire_time) values(?, ? , DATE_ADD(sysdate(),INTERVAL 30 DAY))", userList[0].Iduser, token)
		if err != nil {
			return nil, errors.Wrap(err, "error insert token")
		} else {
			err := repo.Conn.Get(&tokenTest, "select user_id, token, expire_time from session where expire_time >= sysdate() and user_id = ?", userList[0].Iduser)
			if err != nil {
				return nil, errors.Wrap(err, "error select token")
			}
		}
	}

	return tokenTest, nil
}
