/*
  author='du'
  date='2020/2/12 17:26'
*/
package dbops

import (
	"database/sql"
	"log"
	"media/defs"
	"media/util"
	"time"
)

func AddUser(name string, pwd string) error {
	ststIns, err := dbConn.Prepare("insert into users(login_name,pwd) value (?,?)")
	if err != nil {
		return err
	}
	_, err = ststIns.Exec(name, pwd)
	if err != nil {
		return err
	}
	defer ststIns.Close()
	return nil
}

func GetUserCredential(name string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(name).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s", err)
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(name string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name=? and pwd=?")
	if err != nil {
		log.Printf("%s", err)
	}
	_, err = stmtDel.Exec(name, pwd)
	if err != nil {
		log.Printf("%s", err)
	}
	defer stmtDel.Close()
	return nil
}

func GetUser(name string) (*defs.User, error) {
	stmtOut, err := dbConn.Prepare("select id,pwd from users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
	}
	var id int
	var pwd string
	err = stmtOut.QueryRow(name).Scan(&id, &pwd)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &defs.User{Id: id, LoginName: name, Pwd: pwd}
	defer stmtOut.Close()
	return res, nil
}

func AddVideo(userid int, name string) (*defs.VideoInfo, error) {
	vid, err := util.NewUUID()
	if err != nil {
		log.Printf("生成视频Id错误:%s", err)
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006,15:04:05") //这里是golang语言的一个彩蛋。
	stmtIns, err := dbConn.Prepare(`insert into video_info (id,author_id,name,display_ctime) values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, userid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: userid, Name: name, DisplayCtime: ctime}
	defer dbConn.Close()
	return res, nil
}

func GetVideoByVid(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("select * from video_info where vid=?")
	if err != nil {
		log.Printf("获取视频信息出错%s", err)
	}
	var aid int
	var name string
	var ctime string
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer dbConn.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	return res, nil
}

func DeleteVideo(vid string) error {
	stmtDel, err := dbConn.Prepare("delete from video_info where id=?")
	if err != nil {
		log.Printf("删除视频信息失败:%s", err)
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return nil
	}
	defer dbConn.Close()
	return nil
}

func AddComment(aid int, vid string, content string) error {
	id, err := util.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("insert into comments(id,author_id,video_id,content) value (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, aid, vid, content)
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}
