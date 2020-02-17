/*
  author='du'
  date='2020/2/15 21:47'
*/
package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"media/api/dbops"
	"media/api/defs"
	"media/api/response"
	"media/api/session"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "success")
}

//创建用户，注意参数r是指针类型。
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	user := &defs.UserCredential{}

	//序列化判断
	if err := json.Unmarshal(res, user); err != nil {
		response.SendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	//写入db判断
	if err := dbops.AddUser(user.UserName, user.Pwd); err != nil {
		response.SendErrorResponse(w, defs.ErrorDBError)
		return
	}

	//正常情况
	id := session.GenerateSessionId(user.UserName)
	signup := &defs.SignedUp{Success: true, SessionId: id}
	if resp, err := json.Marshal(signup); err != nil {
		response.SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		response.SendNormalResponse(w, string(resp), 201)
	}
}
