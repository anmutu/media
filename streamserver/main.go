/*
  author='du'
  date='2020/2/17 15:08'
*/
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}

//结构体
type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

//hijack
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//panic("implement me")
	if !m.l.GetConn() {
		SendErrorResponse(w, http.StatusTooManyRequests, "请求太多")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.RealseConn()
}

//构造函数
func NewMiddleWareHandler(r *httprouter.Router, count int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(count)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video/:vid", StreamHandler)
	//router.POST("/upload/:vid",UploadHandler)
	return router
}
