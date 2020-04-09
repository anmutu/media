/*
  author='du'
  date='2020/2/12 18:06'
*/
package defs

type UserCredential struct {
	UserName string `json:login_name`
	Pwd      string `json:pwd`
}

type User struct {
	Id        int
	LoginName string
	Pwd       string
}

type Video struct {
	UserId int
	Name   string
}

type VideoInfo struct {
	Id           string `json:"id"`
	AuthorId     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct {
	Id      string `json:"id"`
	VideoId string `json:"video_id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type SimpleSession struct {
	UserName string
	TTL      int64
}

type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}
