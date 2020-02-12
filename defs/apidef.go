/*
  author='du'
  date='2020/2/12 18:06'
*/
package defs


type UserCredential struct {
   UserName string `json:user_name`
   Pwd string `json:pwd`
}

type User struct {
	Id int
	LoginName string
	Pwd string
}
