// the User data access.
package model

import (
	//	"database/sql"
	//"fmt"
	"time"

	"app/logs"
)

type User struct {
	Id         uint64
	UserName   string
	Password   string
	NickName   string
	FullName   string
	Mobile     string
	Email      string
	Role       string
	Status     string
	CreatedOn  time.Time
	ModifiedOn time.Time
}

func (u *User) CheckLogin() (bool, error) {
	query := "SELECT * FROM user_CheckLogin(username=$1, password=$2)"

	res, err := dataStore1.Read(query, u.UserName, u.Password)
	defer CloseResult(res, err, query, u)
	if err != nil {
		logs.Logger.Trace("Error: Failed Login. User:Password: ", u.UserName, u.Password)
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
	//logs.Logger.Trace("Inside user select.", dataStore1)
	query := "SELECT * FROM user_Select(ioffset:=$1,ilimit:=$2)"

	res, err := dataStore1.Read(query, offset, length)
	defer CloseResult(res, err, query, u)
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0, 13)

	var user *User

	//logs.Logger.Trace("Scanning Rows.")
	for res.Next() {
		user = &User{}
		err = res.Scan(
			&user.Id,
			&user.UserName,
			&user.Password,
			&user.NickName,
			&user.FullName,
			&user.Mobile,
			&user.Email,
			&user.Role,
			&user.Status,
			&user.CreatedOn,
			&user.ModifiedOn,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	//logs.Logger.Trace("Size of select: ", len(users))
	return users, nil
}

func (u *User) Delete() (bool, error) {
	query := "SELECT user_Delete(id:=$1)"

	success, err := dataStore1.Modify(query, u.Id)
	if err != nil {
		return false, err
	}

	return success, nil
}

func (u *User) Update() (bool, error) {
	query := "SELECT user_Update(id:=$1,username:=$2,password:=$3,nickname:=$4,fullname:=$5,mobile:=$6,email:=$7,role:=$8,status:=$9"

	success, err := dataStore1.Modify(query, u.Id, u.UserName, u.Password, u.NickName, u.FullName, u.Mobile, u.Email, u.Role, u.Status)
	if err != nil {
		return false, err
	}

	return success, nil
}
