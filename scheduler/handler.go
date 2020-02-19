/*
  author='du'
  date='2020/2/18 2:06'
*/
package main

import (
	"github.com/julienschmidt/httprouter"
	"media/scheduler/s_dbops"
	"net/http"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")

	if len(vid) == 0 {
		sendResponse(w, 400, "video id should not be empty")
		return
	}

	err := s_dbops.AddVideoDelRec(vid)
	if err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}

	sendResponse(w, 200, "")
	return
}
