/*
  author='du'
  date='2020/2/15 17:39'
*/
package session

import "sync"

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDb() {

}

func GenerateSessionId(username string) string {

}

func IsSessionExpired(sid string) (string bool) {

}
