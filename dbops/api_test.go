/*
  author='du'
  date='2020/2/12 19:04'
*/
package dbops

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	err := AddUser("du", "123")
	if err != nil {
		t.Errorf("增加用户出错：%s", err)
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("du", "123")
	if err != nil {
		t.Errorf("删除用户出错:%s", err)
	}

}
