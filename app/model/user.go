package model

import (
	"errors"
)

// Entity is the golang structure for table user.
type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT" json:"id"`                // 用户ID
	Passport string `gorm:"passport;unique_index"   json:"passport"` // 用户账号
	Password string `gorm:"password;size:100"   json:"-"`            // 用户密码
	Nickname string `gorm:"nickname;size:10"   json:"nickname"`      // 用户昵称
	Avatar   string `gorm:"avatar;szie:255"     json:"avatar"`       // 用户头像
	Model
}

func (u *User) Save() (*User, error) {
	db := GetDB()
	db.Create(&u)
	return u, nil
}

func (u *User) Update(reqMap map[string]interface{}) (*User, error) {
	db := GetDB()
	db.Model(u).Updates(reqMap)
	return u, nil
}

func (u *User) GetUserInfoByPassport(passport string) *User {
	var user User
	db := GetDB()
	db.Where("passport = ?", passport).First(&user)
	return &user
}

type UserRoles struct {
	User
	Roles []string `json:"roles"`
}

func (u *User) GetUserInfoById(id uint) (*UserRoles, error) {
	var user UserRoles
	db := GetDB()
	db.Table("user").Where("id = ?", id).Scan(&user)
	if user.Passport != "" {
		user.Roles = []string{"admin", "editor"}
		return &user, nil
	}
	return nil, errors.New("用户不存在")
}

// 检查是否是学校管理员
func CheckSchoolAdmin(school_id uint, user_id uint) bool {
	var user SchoolUser
	db := GetDB()
	db.Table("school_user").Where("school_id = ?", school_id).Where("user_id = ?", user_id).Where("identity = ?", "school").First(&user)
	if user.Id == user_id {
		return true
	}
	return false
}

// 获取所在校区id
func GetCampusIdList(school_id uint, user_id uint) (campusIdList []uint) {
	var schoolUsers []*SchoolUser
	db := GetDB()
	db.Table("school_user").Where("school_id = ?", school_id).Where("user_id = ?", user_id).Where("identity = ?", "teacher").Find(&schoolUsers)
	for _, item := range schoolUsers {
		campusIdList = append(campusIdList, item.CampusId)
	}
	return campusIdList
}
