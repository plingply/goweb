/*
 * @Author: 彭林
 * @Date: 2020-12-25 11:46:41
 * @LastEditTime: 2020-12-25 13:27:17
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/service/course/course.go
 */
package course

import (
	"errors"
	"goframe-web/app/model"
)

func GetCourseList(school_id, campus_id, class_id, page, limit uint) (courseList []*model.CourseList, total uint, msg error) {

	if school_id <= 0 {
		return nil, 0, errors.New("参数错误")
	}

	if campus_id <= 0 {
		return nil, 0, errors.New("参数错误")
	}

	db := model.GetDB()
	db = db.Table("course").Select("course.*, subject.subject_name").Where("course.school_id = ?", school_id).Where("course.campus_id = ?", campus_id)

	if class_id > 0 {
		db = db.Where("course.class_id = ?", class_id)
		db = db.Joins("left join classs on course.class_id = classs.id")
	}

	db = db.Joins("left join subject on course.subject_id = subject.id").Count(&total)

	db.Offset((page - 1) * limit).Limit(limit).Find(&courseList)

	return
}
