package models

import (
	"fmt"
	"ielts/dao"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	//Account string `json:"account" gorm:"not null;unique_index"`
	Account string `json:"account" form:"account"`
	Psw     string `json:"psw" form:"psw"`
	Status  string `json:"status" form:"status" gorm:"default:'no'"`
	Name    string `json:"name" form:"name" gorm:"default:'no'"`
	Time    string `json:"time" form:"time" gorm:"default:'no'"`
	Target  string `json:"target" form:"target" gorm:"default:'no'"`
	Type    string `json:"type" form:"type" gorm:"default:'no'"`
	Avatar  string `json:"avatar" form:"avatar" gorm:"default:'no'"`
}

func RegisterUser(user *User) (code int, user1 *User, err error) {
	account := user.Account
	r := dao.DB.Debug().Where("account = ?", account).First(&User{}).RowsAffected
	if r > 0 {
		fmt.Println("用户已存在")
		return 3001, nil, err
	}
	err = dao.DB.Debug().Create(&user).Error
	if err != nil {
		fmt.Println("创建出错")
		return 3002, nil, err
	}
	user1 = new(User)
	dao.DB.Debug().Where("account = ?", account).First(&user1)
	fmt.Println("内容：：", user1)
	return 3000, user1, err
}
func UpdateUser(u *User) (user *User, err error) {

	var u1 User
	dao.DB.Debug().Where("id = ?", u.ID).First(&u1)

	err = dao.DB.Model(&u1).Debug().Updates(User{Avatar: u.Avatar, Name: u.Name, Time: u.Time, Target: u.Target, Type: u.Type, Status: u.Status}).Error

	if err != nil {
		return nil, err
	}

	user = &u1
	return user, err

}

func UploadAvatar(id int, address string) error {

	err := dao.DB.Debug().Where("id = ?", id).Update("avatar", address).Error
	return err
}
func GEtAUser(account string, psw string) (u *User) {

	var user User
	err := dao.DB.Debug().Where("account = ? AND psw = ? ", account, psw).First(&user).Error
	u = &user
	if err == nil {
		return
	} else {
		return u
	}

}
