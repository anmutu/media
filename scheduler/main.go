/*
  author='du'
  date='2020/2/18 2:05'
*/
package main

import (
	"github.com/julienschmidt/httprouter"
	"media/scheduler/taskrunner"
	"net/http"
)

func main() {
	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}
