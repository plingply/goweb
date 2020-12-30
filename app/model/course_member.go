/*
 * @Author: 彭林
 * @Date: 2020-12-30 11:00:07
 * @LastEditTime: 2020-12-30 11:09:00
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/model/course_member.go
 */
package model

type CourseMember struct {
	Id        uint `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId  uint `gorm:"school_id" json:"school_id"` // 学校ID
	CampusId  uint `gorm:"campus_id" json:"campus_id"` // 学校ID
	CourseId  uint `gorm:"course_id" json:"course_id"`
	StudentId uint `gorm:"student_id" json:"student_id"`
	Model
}

func (c *CourseMember) GetCourseMember(course_id uint) []*CourseMember {
	var members []*CourseMember
	db := GetDB()
	db.Where("course_id = ?", course_id).Find(&members)
	return members
}

func (c *CourseMember) CreateMember(member *CourseMember) bool {
	db := GetDB()
	db.Create(&member)
	return true
}
