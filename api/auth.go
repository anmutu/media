/*
  author='du'
  date='2020/2/15 21:19'
*/
package api

import (
	"media/api/session"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var Header_FIELD_UNAME = "X-User_Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	username, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_SESSION, username)
	return true
}
