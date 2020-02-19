/*
  author='du'
  date='2020/2/18 2:06'
*/
package main

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, sc int, resp string) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
