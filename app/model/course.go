/*
 * @Author: 彭林
 * @Date: 2020-12-25 11:16:42
 * @LastEditTime: 2020-12-29 18:37:43
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/model/course.go
 */
package model

type Course struct {
	Id          uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId    uint   `gorm:"school_id" json:"school_id"` // 学校ID
	CampusId    uint   `gorm:"campus_id" json:"campus_id"` // 学校ID
	SubjectId   string `gorm:"subject_id" json:"subject_id"`
	ClassId     string `gorm:"class_id" json:"class_id"`
	TeacherId   uint   `gorm:"teacher_id" json:"teacher_id"`
	ClassroomId string `gorm:"classroom_id" json:"classroom_id"`
	CourseType  string `gorm:"course_type;size:1"   json:"course_type"` // 1 班课 2学员课
	StartTime   uint   `gorm:"start_time;default:0;type:BIGINT"   json:"start_time"`
	EndTime     uint   `gorm:"end_time;default:0;type:BIGINT"   json:"end_time"`
	Len         uint   `gorm:"len"   json:"len"` // 时长
	Model
}

type CourseList struct {
	Id            uint   `json:"id"`
	SchoolId      uint   `json:"school_id"` // 学校ID
	CampusId      uint   `json:"campus_id"` // 学校ID
	SubjectId     string `json:"subject_id"`
	SubjectName   string `json:"subject_name"`
	ClassId       uint   `json:"class_id"`
	ClassName     string `json:"class_name"`
	TeacherId     uint   `json:"teacher_id"`
	TeacherName   string `json:"teacher_name"`
	ClassroomId   uint   `json:"classroom_id"`
	ClassroomName string `json:"classname_name"`
	CourseType    string `json:"course_type"` // 1 班课 2学员课
	StartTime     uint   `json:"start_time"`
	EndTime       uint   `json:"end_time"`
	Len           uint   `json:"len"`    // 时长
	Status        uint   `json:"status"` // 1 未上课 2 上课中 3已下课
	Model
}

type PaikeParam struct {
	SchoolID    uint   `json:"school_id"`
	CampusID    uint   `json:"campus_id"`
	ClassID     uint   `json:"class_id"`
	ClassroomID uint   `json:"classroom_id"`
	TeacherID   uint   `json:"teacher_id"`
	Len         uint   `json:"len"`
	StartTime   uint   `json:"start_time"`
	Type        uint   `json:"type"`
	Conflict    string `json:"conflict"`
}

func (c *Course) CheckCourse(course *PaikeParam) *PaikeParam {

	db := GetDB()
	endTime := course.StartTime + course.Len*60000

	if course.ClassID > 0 {
		var cx []Course
		db.Table("course").Where("start_time < ? && end_time >= ?", course.StartTime, course.StartTime).Or("start_time >= ? && start_time <= ?", course.StartTime, endTime).Find(&cx)
		if len(cx) > 0 {
			course.Conflict = "class"
			return course
		}
	}

	// SELECT * FROM `course` WHERE `course`.`deleted_at` IS NULL AND (class_id = 1) AND ((start_time < 1608187186600 AND end_time >= 1608187186600) OR (start_time >= 1608187186600 AND start_time <= 1608189886600))

	// if course.TeacherID > 0 {
	// 	var cx []Course
	// 	db.Raw("SELECT * FROM course WHERE course.deleted_at IS NULL && teacher_id = ? && ((start_time < ? && end_time >= ?) || (start_time >= ? && start_time <= ?))", course.TeacherID, course.StartTime, course.StartTime, course.StartTime, endTime).Find(&cx)
	// 	if len(cx) > 0 {
	// 		course.Conflict = "teacher"
	// 		return course
	// 	}
	// }

	// if course.ClassroomID > 0 {
	// 	var cx []Course
	// 	db.Raw("SELECT * FROM course WHERE course.deleted_at IS NULL && classroom_id = ? && ((start_time < ? && end_time >= ?) || (start_time >= ? && start_time <= ?))", course.ClassroomID, course.StartTime, course.StartTime, course.StartTime, endTime).Find(&cx)
	// 	if len(cx) > 0 {
	// 		course.Conflict = "classroom"
	// 		return course
	// 	}
	// }

	return course
}
