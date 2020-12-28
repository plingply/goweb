/*
 * @Author: 彭林
 * @Date: 2020-12-25 11:16:42
 * @LastEditTime: 2020-12-28 17:36:32
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/model/course.go
 */
package model

type Course struct {
	Id         uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId   uint   `gorm:"school_id" json:"school_id"` // 学校ID
	CampusId   uint   `gorm:"campus_id" json:"campus_id"` // 学校ID
	SubjectId  string `gorm:"subject_id" json:"subject_id"`
	ClassId    string `gorm:"class_id" json:"class_id"`
	CourseType string `gorm:"course_type;size:1"   json:"course_type"` // 1 班课 2学员课
	StartTime  uint   `gorm:"start_time;default:0;type:BIGINT"   json:"start_time"`
	EndTime    uint   `gorm:"end_time;default:0;type:BIGINT"   json:"end_time"`
	Len        uint   `gorm:"len"   json:"len"` // 时长
	Model
}

type CourseList struct {
	Id          uint   `json:"id"`
	SchoolId    uint   `json:"school_id"` // 学校ID
	CampusId    uint   `json:"campus_id"` // 学校ID
	SubjectId   string `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	ClassId     string `json:"class_id"`
	ClassName   string `json:"class_name"`
	CourseType  string `json:"course_type"` // 1 班课 2学员课
	StartTime   uint   `json:"start_time"`
	EndTime     uint   `json:"end_time"`
	Len         uint   `json:"len"`    // 时长
	Status      uint   `json:"status"` // 1 未上课 2 上课中 3已下课
	Model
}
