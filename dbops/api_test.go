/*
  author='du'
  date='2020/2/12 19:04'
*/
package dbops

import (
	"testing"
)

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func clearTables() {
	dbConn.Exec("truncate users")
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
