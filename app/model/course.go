/*
 * @Author: 彭林
 * @Date: 2020-12-25 11:16:42
 * @LastEditTime: 2020-12-30 11:20:18
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/model/course.go
 */
package model

import "errors"

type Course struct {
	Id          uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId    uint   `gorm:"school_id" json:"school_id"` // 学校ID
	CampusId    uint   `gorm:"campus_id" json:"campus_id"` // 学校ID
	SubjectId   uint   `gorm:"subject_id" json:"subject_id"`
	ClassId     uint   `gorm:"class_id" json:"class_id"`
	TeacherId   uint   `gorm:"teacher_id" json:"teacher_id"`
	ClassroomId uint   `gorm:"classroom_id" json:"classroom_id"`
	CourseType  uint   `gorm:"course_type;size:1"   json:"course_type"` // 1 班课 2学员课
	StartTime   uint   `gorm:"start_time;default:0;type:BIGINT"   json:"start_time"`
	EndTime     uint   `gorm:"end_time;default:0;type:BIGINT"   json:"end_time"`
	Len         uint   `gorm:"len"   json:"len"`           // 时长
	Note        string `gorm:"len;size:200"   json:"note"` // 时长
	Model
}

type CourseList struct {
	Id            uint   `json:"id"`
	SchoolId      uint   `json:"school_id"` // 学校ID
	CampusId      uint   `json:"campus_id"` // 学校ID
	SubjectId     uint   `json:"subject_id"`
	SubjectName   string `json:"subject_name"`
	ClassId       uint   `json:"class_id"`
	ClassName     string `json:"class_name"`
	TeacherId     uint   `json:"teacher_id"`
	TeacherName   string `json:"teacher_name"`
	ClassroomId   uint   `json:"classroom_id"`
	ClassroomName string `json:"classname_name"`
	CourseType    uint   `json:"course_type"` // 1 班课 2学员课
	StartTime     uint   `json:"start_time"`
	EndTime       uint   `json:"end_time"`
	Len           uint   `json:"len"`    // 时长
	Status        uint   `json:"status"` // 1 未上课 2 上课中 3已下课
	Note          string `json:"note"`   // 时长
	Model
}

type PaikeParam struct {
	SchoolID    uint   `json:"school_id"`
	CampusID    uint   `json:"campus_id"`
	ClassID     uint   `json:"class_id"`
	ClassroomID uint   `json:"classroom_id"`
	TeacherID   uint   `json:"teacher_id"`
	SubjectID   uint   `json:"subject_id"`
	Len         uint   `json:"len"`
	StartTime   uint   `json:"start_time"`
	Type        uint   `json:"type"`
	Note        string `json:"note"` // 时长
	Conflict    string `json:"conflict"`
}

func (c *Course) CheckCourse(course *PaikeParam) *PaikeParam {

	db := GetDB()
	endTime := course.StartTime + course.Len*60000
	db = db.Table("course").Where("start_time < ? && end_time >= ?", course.StartTime, course.StartTime).Or("start_time >= ? && start_time <= ?", course.StartTime, endTime)

	if course.ClassID > 0 {
		var cx []Course
		db.Where("class_id = ?", course.ClassID).Find(&cx)
		if len(cx) > 0 {
			course.Conflict = "class"
			return course
		}
	}

	if course.TeacherID > 0 {
		var cx []Course
		db.Where("teacher_id = ?", course.TeacherID).Find(&cx)
		if len(cx) > 0 {
			course.Conflict = "teacher"
			return course
		}
	}

	if course.ClassroomID > 0 {
		var cx []Course
		db.Where("classroom_id = ?", course.ClassroomID).Find(&cx)
		if len(cx) > 0 {
			course.Conflict = "classroom"
			return course
		}
	}

	return course
}

func (c *Course) AddCourse(course *PaikeParam) error {
	var cmodel Course
	cmodel.SchoolId = course.SchoolID
	cmodel.CampusId = course.CampusID
	cmodel.ClassId = course.ClassID
	cmodel.TeacherId = course.TeacherID
	cmodel.SubjectId = course.SubjectID
	cmodel.ClassroomId = course.ClassroomID
	cmodel.CourseType = 1
	cmodel.StartTime = course.StartTime
	cmodel.Len = course.Len
	cmodel.EndTime = course.StartTime + cmodel.Len*60000

	db := GetDB()
	db.Create(&cmodel)

	if cmodel.Id != 0 {
		// 班课
		if cmodel.CourseType == 1 {
			go func(id uint) {
				var classMember ClassMember
				memberList, _ := classMember.GetClassMemberList(cmodel.ClassId, 1, 0, 0)
				for _, v := range memberList {
					var courseMember CourseMember
					courseMember.CampusId = cmodel.CampusId
					courseMember.SchoolId = cmodel.SchoolId
					courseMember.StudentId = v.StudentId
					courseMember.CourseId = id
					courseMember.CreateMember(&courseMember)
				}
			}(cmodel.Id)
		}

		return nil
	} else {
		return errors.New("创建失败")
	}

}
