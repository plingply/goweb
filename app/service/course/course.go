/*
 * @Author: 彭林
 * @Date: 2020-12-25 11:46:41
 * @LastEditTime: 2020-12-30 10:54:57
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/service/course/course.go
 */
package course

import (
	"errors"
	"goframe-web/app/model"
	"sync"

	"github.com/gogf/gf/os/gtime"
)

func GetCourseList(school_id, campus_id, class_id, start_time, end_time, page, limit uint) (courseList []*model.CourseList, total uint, msg error) {

	if school_id <= 0 {
		return nil, 0, errors.New("参数错误")
	}

	if campus_id <= 0 {
		return nil, 0, errors.New("参数错误")
	}

	db := model.GetDB()
	db = db.Table("course").Select("course.*, subject.subject_name").Where("course.school_id = ?", school_id).Where("course.campus_id = ?", campus_id)

	if start_time > 0 && end_time > 0 && end_time > start_time {
		db = db.Where("course.start_time >= ?", start_time)
		db = db.Where("course.start_time <= ?", end_time)
	}

	if class_id > 0 {
		db = db.Where("course.class_id = ?", class_id)
		db = db.Joins("left join classs on course.class_id = classs.id")
	}

	db = db.Joins("left join subject on course.subject_id = subject.id").Count(&total)

	if page > 0 && limit > 0 {
		db = db.Offset((page - 1) * limit).Limit(limit)
	}

	db.Find(&courseList)

	for i, v := range courseList {
		start := v.StartTime
		end := v.EndTime
		now := uint(gtime.TimestampMilli())

		if start > now {
			courseList[i].Status = 1
		}

		if start < now && end >= now {
			courseList[i].Status = 2
		}

		if end < now {
			courseList[i].Status = 3
		}
	}

	return
}

func CheckCourse(courseList []*model.PaikeParam) []*model.PaikeParam {
	var course model.Course
	var wg sync.WaitGroup
	for _, v := range courseList {
		wg.Add(1)
		go func(v *model.PaikeParam) {
			course.CheckCourse(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	return courseList
}

func AddCourse(courseList []*model.PaikeParam) (uint, error) {
	var course model.Course
	var count uint = 0
	var wg sync.WaitGroup
	for _, v := range courseList {
		count++
		wg.Add(1)
		go func(v *model.PaikeParam) {
			course.AddCourse(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	return count, nil
}
