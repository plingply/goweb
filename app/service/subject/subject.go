package subject

import (
	"errors"
	"goframe-web/app/model"
	"unicode/utf8"
)

func GetSubjectList(school_id, campus_id, page, limit uint) (result interface{}, total int, err error) {

	if school_id == 0 {
		return nil, 0, errors.New("学校id不能为空")
	}

	if campus_id == 0 {
		return nil, 0, errors.New("校区id不能为空")
	}

	var subject *model.Subject
	result, total = subject.GetSubjectList(school_id, campus_id, page, limit)

	return result, total, nil
}

func UpdateSubject(subject_id uint, data map[string]interface{}) (re bool, msg error) {

	if subject_id <= 0 {
		return false, errors.New("参数错误")
	}

	if data["subject_name"] != nil && utf8.RuneCountInString(data["subject_name"].(string)) > 20 {
		return false, errors.New("参数错误 subject_name")
	}

	if data["remark"] != nil && utf8.RuneCountInString(data["remark"].(string)) > 100 {
		return false, errors.New("参数错误 remark")
	}

	var subject *model.Subject
	re = subject.UpdateSubject(subject_id, data)

	return
}

func CreateSubject(school_id, campus_id uint, subject *model.Subject) (re bool, msg error) {

	if school_id <= 0 {
		return false, errors.New("参数错误")
	}

	if campus_id <= 0 {
		return false, errors.New("参数错误")
	}

	var subjectModel *model.Subject
	re = subjectModel.CreateSubject(subject)

	return
}
