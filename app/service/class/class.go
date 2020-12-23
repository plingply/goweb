package class

import (
	"errors"
	"goframe-web/app/model"
	"unicode/utf8"
)

func GetClassList(school_id, campus_id, user_id, page, limit uint) (result interface{}, total int, err error) {

	if school_id == 0 {
		return nil, 0, errors.New("学校id不能为空")
	}
	if page <= 0 {
		return nil, 0, errors.New("参数错误")
	}
	if limit <= 0 {
		return nil, 0, errors.New("参数错误")
	}

	var class *model.Classs
	result, total = class.GetClassList(school_id, campus_id, user_id, page, limit)

	return result, total, nil
}

func GetClassSimpleList(school_id, campus_id, user_id uint) (result interface{}, err error) {
	if school_id == 0 {
		return nil, errors.New("学校id不能为空")
	}
	var class *model.Classs
	result = class.GetClassSimpleList(school_id, campus_id, user_id)
	return result, nil
}

func UpdateClass(class_id uint, data map[string]interface{}) (re bool, msg error) {

	if class_id <= 0 {
		return false, errors.New("参数错误")
	}

	if data["class_name"] != nil && utf8.RuneCountInString(data["class_name"].(string)) > 20 {
		return false, errors.New("参数错误 Class_name")
	}

	var class *model.Classs
	re = class.UpdateClass(class_id, data)

	return
}

func CreateClass(school_id, campus_id uint, class *model.Classs) (re bool, msg error) {

	if school_id <= 0 {
		return false, errors.New("参数错误")
	}

	if campus_id <= 0 {
		return false, errors.New("参数错误")
	}

	var classModel *model.Classs
	re = classModel.CreateClass(class)

	return
}
