/*
  author='du'
  date='2020/2/17 15:24'
*/
package main

import (
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}
