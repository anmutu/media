/*
  author='du'
  date='2020/2/12 17:26'
*/
package dbops

import (
	"database/sql"
	"log"
	"media/defs"
)

func AddUser(name string, pwd string) error {
	ststIns, err := dbConn.Prepare("insert into users(login_name,pwd) value (?,?)")
	if err != nil {
		return err
	}
	_, err = ststIns.Exec(name, pwd)
	if err != nil {
		return err
	}
	defer ststIns.Close()
	return nil
}

func GetUserCredential(name string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(name).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s", err)
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(name string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where user=? and pwd=?")
	if err != nil {
		log.Printf("%s", err)
	}
	_, err = stmtDel.Exec(name, pwd)
	if err != nil {
		log.Printf("%s", err)
	}
	defer stmtDel.Close()
	return nil
}

func GetUser(name string) (*defs.User, error) {
	stmtOut, err := dbConn.Prepare("select id,pwd from users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
	}
	var id int
	var pwd string
	err = stmtOut.QueryRow(name).Scan(&id, &pwd)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &defs.User{Id: id, LoginName: name, Pwd: pwd}
	defer stmtOut.Close()
	return res, nil
}
