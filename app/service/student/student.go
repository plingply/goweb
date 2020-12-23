package student

import (
	"errors"
	"goframe-web/app/model"
	"unicode/utf8"
)

func GetStudentList(school_id, campus_id, page, limit uint) (result interface{}, total int, err error) {

	if school_id == 0 {
		return nil, 0, errors.New("学校id不能为空")
	}

	if campus_id == 0 {
		return nil, 0, errors.New("校区id不能为空")
	}

	if page <= 0 {
		return nil, 0, errors.New("参数错误")
	}
	if limit <= 0 {
		return nil, 0, errors.New("参数错误")
	}

	var student *model.Student
	result, total = student.GetStudentList(school_id, campus_id, page, limit)

	return result, total, nil
}

func UpdateStudent(student_id uint, data map[string]interface{}) (re bool, msg error) {

	if student_id <= 0 {
		return false, errors.New("参数错误")
	}

	if utf8.RuneCountInString(data["student_name"].(string)) > 20 {
		return false, errors.New("参数错误 student_name")
	}

	if utf8.RuneCountInString(data["avatar"].(string)) > 100 {
		return false, errors.New("参数错误 avatar")
	}

	if utf8.RuneCountInString(data["address"].(string)) > 50 {
		return false, errors.New("参数错误 address")
	}

	if utf8.RuneCountInString(data["school_name"].(string)) > 50 {
		return false, errors.New("参数错误 school_name")
	}

	if utf8.RuneCountInString(data["birthday"].(string)) > 20 {
		return false, errors.New("参数错误 birthday")
	}

	if utf8.RuneCountInString(data["remark"].(string)) > 100 {
		return false, errors.New("参数错误 remark")
	}

	var student *model.Student
	re = student.UpdateStudent(student_id, data)

	return
}

func CreateStudent(school_id, campus_id uint, student *model.Student) (re bool, msg error) {

	if school_id <= 0 {
		return false, errors.New("参数错误")
	}

	if campus_id <= 0 {
		return false, errors.New("参数错误")
	}

	var studentModel *model.Student
	re = studentModel.CreateStudent(student)

	return
}
