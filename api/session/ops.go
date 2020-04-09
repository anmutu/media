/*
  author='du'
  date='2020/2/15 17:39'
*/
package session

import (
	"media/api/dbops"
	"media/api/defs"
	"media/api/util"
	"sync"
	"time"
)

//线程安全的一个map,这个是1.9以后才有的。
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowElaspedSecond() int64 {
	return time.Now().UnixNano() / 1000000
}

//获取所有session
func LoadSessionFromDb() {
	sessions, err := dbops.GetAllSession()
	if err != nil {
		return
	}
	sessions.Range(func(k, value interface{}) bool {
		session := value.(*defs.SimpleSession)
		sessionMap.Store(k, session)
		return true
	})
}

//生成session
func GenerateSessionId(username string) string {
	id, _ := util.NewUUID()
	t := nowElaspedSecond()
	ttl := t + 60*60*100 //暂时设置为1个小时
	session := &defs.SimpleSession{UserName: username, TTL: ttl}
	sessionMap.Store(id, session)
	dbops.InsertSession(id, ttl, username)
	return id
}

//是否过期，true为过期，false为未过期。
func IsSessionExpired(sid string) (string, bool) {
	session, ok := sessionMap.Load(sid)
	if ok {
		t := nowElaspedSecond()
		if session.(*defs.SimpleSession).TTL < t {
			return "", true //表示已过期
		} else {
			return session.(*defs.SimpleSession).UserName, false
		}
	}
	return "", true
}

//删除过期session
func deleteExpiredSession(sid string) error {
	sessionMap.Delete(sid)
	if err := dbops.DeleteSession(sid); err != nil {
		return err
	} else {
		return nil
	}
}
