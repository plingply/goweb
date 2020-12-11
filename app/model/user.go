package model

import (
	"errors"
)

// Entity is the golang structure for table user.
type User struct {
	Id       uint        `gorm:"AUTO_INCREMENT" json:"id"`        // 用户ID
	Passport string      `gorm:"passport;unique_index"   json:"passport"`  // 用户账号
	Password string      `gorm:"password"   json:"-"`  // 用户密码
	Nickname string      `gorm:"nickname"   json:"nickname"`  // 用户昵称
	Model
}

func (u *User)Save() (*User, error)  {
	db := GetDB()
	db.Create(&u)
	return u, nil
}

func (u *User)Update(reqMap map[string]interface{})(*User, error){
	db := GetDB()
	db.Model(u).Updates(reqMap)
	return u, nil
}

func (u *User)GetUserInfoByPassport(passport string) *User  {
	var user User
	db := GetDB()
	db.Where("passport = ?", passport).First(&user)
	return &user
}

func (u *User)GetUserInfoById(id uint) (*User, error)  {
	var user User
	user.Id = id
	db := GetDB()
	db.First(&user)
	if user.Passport != "" {
		return &user, nil
	}
	return nil, errors.New("用户不存在")
}