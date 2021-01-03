package school

import (
	"errors"
	"goframe-web/app/model"
	"goframe-web/app/service/user"
)

func CreateSchool(campus_name string, user_id uint, school *model.School) (msg error) {

	if school.SchoolName == "" {
		return errors.New("参数错误")
	}

	if school.Logo == "" {
		return errors.New("参数错误")
	}

	var schoolModel *model.School
	schools := schoolModel.CreateSchool(school)

	if schools.Id == 0 {
		return errors.New("创建失败")
	} else {
		var campus model.Campus
		campus.SchoolId = schools.Id
		campus.CampusName = campus_name
		if re := campus.CreateCampus(&campus); re {

			userInfo, err := user.GetUserInfo(user_id)

			if err != nil {
				return errors.New("创建失败")
			}
			var teacher model.SchoolUser
			teacher.SchoolId = schools.Id
			teacher.CampusId = 0
			teacher.Identity = "school"
			teacher.UserId = user_id
			teacher.Sex = 3
			teacher.TeacherName = userInfo.Nickname
			teacher.Phone = userInfo.Passport
			if r := teacher.CreateTeacher(&teacher); r {
				return nil
			} else {
				return errors.New("创建失败")
			}
		} else {
			return errors.New("创建失败")
		}
	}
}
