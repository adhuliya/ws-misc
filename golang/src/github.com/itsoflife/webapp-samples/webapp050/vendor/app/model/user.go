// the User data access.
package model

import (
	"database/sql"
	"fmt"
	"time"

	"app/logs"
)

type User struct {
	id         uint64
	username   string
	password   string
	nickname   string
	fullname   string
	mobile     string
	email      string
	role       string
	status     string
	createdOn  time.Time
	modifiedOn time.Time
}

func (u *User) CheckLogin() (bool, err) {
	query := "SELECT * FROM user_CheckLogin(username=$1, password=$2)"

	res, err := dataStore1.Read(query, u.username, u.password)
	defer CloseResult(res, err, query, u)
	if err != nil {
		return false, err
	}

	isCorrect := false
	for res.Next() {
		isCorrect = true
		break
	}

	return isCorrect, nil
}

func (u *User) Select(offset uint64, length uint64) ([]*User, error) {
	query := "SELECT * FROM user_Select(ioffset:=$1,ilimit:=$2)"

	res, err = dataStore1.conn.Read(query, offset, length)
	defer CloseResult(res, err, query, u)
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0, 13)

	var user User

	for res.Next() {
		user = &User{}
		err = res.Scan(
			&user.id,
			&user.username,
			&user.password,
			&user.nickname,
			&user.fullname,
			&user.mobile,
			&user.email,
			&user.role,
			&user.status,
			&user.createdOn,
			&user.modifiedOn,
		)

		if err != nil {
			return nil, err
		}

		users.append(user)
	}

	return users, nil
}

func (u *User) Delete() (bool, error) {
	query := "SELECT user_Delete(id:=$1)"

	res, err := dataStore1.conn.Modify(query, u.id)
	defer CloseResult(res, err, query, u)
	if err != nil {
		return false, err
	}

	b := ScanBool(res)
	return b, nil
}

func (u *User) Update() (bool, error) {
	query := "SELECT user_Update(id:=$1,username:=$2,password:=$3,nickname:=$4,fullname:=$5,mobile:=$6,email:=$7,role:=$8,status:=$9"

	res, err := dataStore1.conn.Modify(query, u.id, u.username, u.password, u.nickname, u.fullname, u.mobile, u.email, u.role, u.status)
	defer CloseResult(res, err, query, u)
	if err != nil {
		return false, err
	}

	b := ScanBool(res)
	return b, nil
}
