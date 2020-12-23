package campus

import (
	"errors"
	"goframe-web/app/model"
	"unicode/utf8"
)

func GetCampusList(schoolId, user_id, page, limit uint) (result interface{}, total int, err error) {
	if schoolId == 0 {
		return nil, 0, errors.New("学校id不能为空")
	}
	if page <= 0 {
		return nil, 0, errors.New("参数错误")
	}
	if limit <= 0 {
		return nil, 0, errors.New("参数错误")
	}
	var campus *model.Campus
	result, total = campus.GetCampusList(schoolId, user_id, page, limit)

	return result, total, nil
}

func GetCampusSimpleList(schoolId, user_id uint) (result interface{}, err error) {
	if schoolId == 0 {
		return nil, errors.New("学校id不能为空")
	}
	var campus *model.Campus
	result, err = campus.GetCampusSimpleList(schoolId, user_id)
	return
}

func UpdateCampus(campusId uint, data map[string]interface{}) (re bool, msg error) {

	if campusId <= 0 {
		return false, errors.New("参数错误")
	}

	if utf8.RuneCountInString(data["campus_name"].(string)) > 20 {
		return false, errors.New("参数错误 campus_name")
	}

	if utf8.RuneCountInString(data["address"].(string)) > 50 {
		return false, errors.New("参数错误 address")
	}

	if utf8.RuneCountInString(data["province"].(string)) > 7 {
		return false, errors.New("参数错误 province")
	}

	if utf8.RuneCountInString(data["city"].(string)) > 7 {
		return false, errors.New("参数错误 city")
	}

	if utf8.RuneCountInString(data["area"].(string)) > 7 {
		return false, errors.New("参数错误 area")
	}

	var campus *model.Campus
	re = campus.UpdateCampus(campusId, data)

	return
}

func CreateCampus(schoolId uint, campus *model.Campus) (re bool, msg error) {

	if schoolId <= 0 {
		return false, errors.New("参数错误")
	}

	var campusModel *model.Campus
	campus.SchoolId = schoolId
	re = campusModel.CreateCampus(campus)

	return
}
