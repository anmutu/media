/*
  author='du'
  date='2020/2/15 17:50'
*/
package dbops

import (
	"database/sql"
	"log"
	"media/api/defs"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, username string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("insert into sessions(session_id,TTL,login_name)values (?,?,?)")
	if err != nil {
		return nil
	}
	_, err = stmtIns.Exec(sid, ttlstr, username)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetSessionById(sid string) (*defs.SimpleSession, error) {
	session := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select TTL,login_name from sessions from where session_id=?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var username string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		session.TTL = res
		session.UserName = username
	} else {
		return nil, err
	}
	defer dbConn.Close()
	return session, nil
}

func GetAllSession() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("select * from sessions")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err = rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("获取session错误：%s", err)
			break
		}
		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			session := &defs.SimpleSession{UserName: login_name, TTL: ttl}
			m.Store(id, session)
		}
	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("select * from sessions where session_id=?")
	if err != nil {
		return err
	}

	if _, err := stmtOut.Query(sid); err != nil {
		return err
	}
	return nil
}
