/*
  author='du'
  date='2020/2/12 19:04'
*/
package dbops

import (
	"testing"
)

var videoid string

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", TestAddUser)
	t.Run("Get", TestGetUserCredential)
	t.Run("Del", TestDeleteUser)
}

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
}

func TestAddUser(t *testing.T) {
	err := AddUser("du", "123")
	if err != nil {
		t.Errorf("增加用户出错：%s", err)
	}
}

func TestGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("du")
	if err != nil && pwd != "123" {
		t.Errorf("获取用户失败:%s", err)
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("du", "123")
	if err != nil {
		t.Errorf("删除用户出错:%s", err)
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddVideo", TestAddVideo)
	t.Run("GetVideo", TestGetVideoByVid)
	t.Run("DelVideo", TestDeleteVideo)
}

func TestAddVideo(t *testing.T) {
	vedio, err := AddVideo(1, "duvideo")
	if err != nil {
		t.Errorf("增加视频信息失败:%s", err)
	}
	videoid = vedio.Id
}

func TestGetVideoByVid(t *testing.T) {
	_, err := GetVideoByVid(videoid)
	if err != nil {
		t.Errorf("获取视频信息失败:%s", err)
	}
}

func TestDeleteVideo(t *testing.T) {
	err := DeleteVideo(videoid)
	if err != nil {
		t.Errorf("删除信息失败：%s", err)
	}
}
