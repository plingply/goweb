package school_user

import (
	"errors"
	"goframe-web/app/model"
	"goframe-web/app/service/user"
)

type SchoolListByUser struct {
	SchoolId   uint   `json:"school_id"` // 学校ID
	SchoolName string `json:"school_name"`
}

func GetSchoolList(user_id uint) []*SchoolListByUser {
	var schoolList []*SchoolListByUser
	db := model.GetDB()
	db.Table("school_user").Select("school_user.school_id, school.school_name as school_name").Group("school_user.school_id").Joins("left join school on school_user.school_id = school.id").Where("user_id = ?", user_id).Scan(&schoolList)
	return schoolList
}

func GetTeacherList(schoolId, campusId, page, Limit uint) (result interface{}, total int, err error) {

	if schoolId == 0 {
		return nil, 0, errors.New("学校id不能为空")
	}
	if campusId == 0 {
		return nil, 0, errors.New("校区id不能为空")
	}
	if page <= 0 {
		return nil, 0, errors.New("参数错误")
	}
	if Limit <= 0 {
		return nil, 0, errors.New("参数错误")
	}
	var schoolUser *model.SchoolUser
	result, total = schoolUser.GetTeacherList(schoolId, campusId, page, Limit)

	return result, total, nil
}

func UpdateTeacher(teacher_id uint, data map[string]interface{}) (re bool, msg error) {

	if teacher_id <= 0 {
		return false, errors.New("参数错误")
	}

	var schoolUser *model.SchoolUser
	re = schoolUser.UpdateTeacher(teacher_id, data)

	return
}

func CreateTeacher(campusId uint, schoolUser *model.SchoolUser) (re bool, msg error) {

	if campusId <= 0 {
		return false, errors.New("参数错误")
	}

	// 账号唯一性数据检查
	if ok := CheckTeacher(campusId, schoolUser.Phone); ok {
		return false, errors.New("老师已经存在")
	}

	// 判断是否已经存在user
	ok, user_id := user.CheckPassport(schoolUser.Phone)
	if ok {
		schoolUser.UserId = user_id
	} else {
		var SignUpParam user.SignUpParam
		SignUpParam.Nickname = schoolUser.TeacherName
		SignUpParam.Passport = schoolUser.Phone
		SignUpParam.Password = "111111"
		SignUpParam.Password2 = "111111"
		result, err := user.SignUp(&SignUpParam)
		if err != nil {
			return false, err
		}
		schoolUser.UserId = result.(*model.User).Id
	}

	var schoolUserModel model.SchoolUser
	re = schoolUserModel.CreateTeacher(schoolUser)

	return
}

// 检查账号是否存在
func CheckTeacher(campus_id uint, phone string) bool {
	var user model.SchoolUser
	teacher := user.GetTeacherInfoByPhone(campus_id, phone)
	if teacher.Id != 0 {
		return true
	}
	return false
}
