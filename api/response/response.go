/*
  author='du'
  date='2020/2/15 21:20'
*/
package response

import (
	"encoding/json"
	"io"
	"media/api/defs"
	"net/http"
)

//错误返回
func SendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

//正常返回
func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
