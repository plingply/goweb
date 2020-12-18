package school_user

import "goframe-web/app/model"

type SchoolListByUser struct {
	SchoolId   uint   `gorm:"school_id" json:"school_id"` // 学校ID
	SchoolName string `gorm:"school_name" json:"school_name"`
}

func GetSchoolList(user_id uint) []*SchoolListByUser {
	var schoolList []*SchoolListByUser
	db := model.GetDB()
	db.Table("school_user").Select("school_user.school_id, school.school_name as school_name").Group("school_user.school_id").Joins("left join school on school_user.school_id = school.id").Where("user_id = ?", user_id).Scan(&schoolList)
	return schoolList
}
