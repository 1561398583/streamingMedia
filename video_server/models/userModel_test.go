package models

import (
	"fmt"
	"testing"
)


func TestAddUser(t *testing.T) {
	name := "yx"
	pwd := "123456"
	err := AddUser(name, pwd)
	if err != nil {
		t.Errorf("error of addUser %v", err)
	}
}

//增加名字已经存在的user
func TestAddSameUser(t *testing.T) {
	name := "yx"
	pwd := "123456"
	err := AddUser(name, pwd)
	if err != nil {
		t.Errorf("error of addUser %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser(56)
	if err != nil {
		t.Errorf("error of addUser %v", err)
	}
}

func TestChangePwd(t *testing.T) {
	err := ChangePwd(uint(56), "zyddlmm123")
	if err != nil {
		t.Errorf("ChangePwd error : %v", err)
	}
}

func TestGetUserById(t *testing.T) {
	user, err := GetUserById(56)
	if err != nil {
		t.Errorf("getUser error : %v", err)
		return
	}
	if user == nil {
		fmt.Println("no this user")
	}else {
		fmt.Printf("find this user : name=%s ; pwd=%s", user.Name, user.Pwd)
	}
}

