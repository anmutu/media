/*
  author='du'
  date='2020/2/15 21:20'
*/
package main

import (
	"github.com/julienschmidt/httprouter"
	"media/api/auth"
	"media/api/handler"
	"media/api/session"
	"net/http"
)

//中间件handler结构体
type middleWareHandler struct {
	r *httprouter.Router
}

//构造函数
func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

//结构体的方法
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//检查session，也就是检查其合法性
	auth.ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/test", handler.Test)
	router.POST("/user", handler.CreateUser)
	return router
}

func Prepare() {
	session.LoadSessionFromDb()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
