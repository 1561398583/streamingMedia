package models

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"video_server/weberrors"
	_ "video_server/weberrors"
)

type User struct {
	gorm.Model
	Name string
	Pwd string
}

func AddUser(name string, pwd string)  error{
	user := User{Name: name, Pwd: pwd}
	result := DB.Create(&user)
	if result.Error != nil {
		if dupErr, ok := result.Error.(*mysql.MySQLError); ok {
			if dupErr.Number == 1062 {	//重复
				return weberrors.New(name + "已存在", 2)
			}
			return result.Error
		}
	}
	return nil
}

func DeleteUser(userId uint)  error{
	result := DB.Delete(&User{}, userId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserById(userId uint)  (*User,error){
	user := User{}
	r := DB.Where("id = ?", userId).First(&user)
	if r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound{	//如果没有这行
			return nil, nil
		}else {		//否则就是真的出错了
			return nil, r.Error
		}
	}
	return &user, nil
}

func GetUserByName(name string)  (*User,error){
	user := User{}
	r := DB.Where("name = ?", name).First(&user)
	if r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound{	//如果没有这行
			return nil, nil
		}else {		//否则就是真的出错了
			return nil, r.Error
		}
	}
	return &user, nil
}

func ChangePwd(userId uint, newPwd string)  error{
	r := DB.Model(&User{}).Where("id = ?", userId).Update("pwd", newPwd)
	if r.Error != nil {
		return r.Error
	}
	return nil
}