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
	http.ListenAndServe(":9000", r)
}

type middleWareHandler struct {
	r *httprouter.Handle
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	//router.GET("/video/:vid",StreamHandler)
	//router.POST("/upload/:vid",UploadHandler)
	return router
}
