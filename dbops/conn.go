/*
  author='du'
  date='2020/2/12 17:28'
*/
package dbops

import (
	"database/sql"
)

var (
	dbConn *sql.DB
	err error
)

func init(){
	dbConn,err=sql.Open("mysql","root:root!@#@tcp(localhost:3306)/video_server?charset=utf8")
	if err!=nil{
		panic(err.Error())
	}
}
