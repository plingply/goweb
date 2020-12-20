package campus

import (
	"errors"
	"goframe-web/app/model"
)

func GetCampusList(schoolId uint, page uint, Limit uint) (result interface{}, total int, err error) {
	if schoolId == 0 {
		return nil, 0, errors.New("学校id不能为空")
	}
	if page <= 0 {
		return nil, 0, errors.New("参数错误")
	}
	if Limit <= 0 {
		return nil, 0, errors.New("参数错误")
	}
	var campus *model.Campus
	result, total = campus.GetCampusList(schoolId, page, Limit)

	return result, total, nil
}

func GetCampusSimpleList(schoolId uint) (result interface{}, err error) {
	if schoolId == 0 {
		return nil, errors.New("学校id不能为空")
	}
	var campus *model.Campus
	result, err = campus.GetCampusSimpleList(schoolId)
	return
}

func UpdateCampus(campusId uint, data map[string]interface{}) (re bool, msg error) {

	if campusId <= 0 {
		return false, errors.New("参数错误")
	}

	if len(data["campus_name"].(string)) > 20 {
		return false, errors.New("参数错误 campus_name")
	}

	if len(data["address"].(string)) > 50 {
		return false, errors.New("参数错误 address")
	}

	if len(data["province"].(string)) > 7 {
		return false, errors.New("参数错误 province")
	}

	if len(data["city"].(string)) > 7 {
		return false, errors.New("参数错误 city")
	}

	if len(data["area"].(string)) > 7 {
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
