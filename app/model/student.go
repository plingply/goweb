/*
 * @Author: 彭林
 * @Date: 2020-12-23 14:13:01
 * @LastEditTime: 2020-12-30 14:46:10
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/model/student.go
 */
package model

type Student struct {
	Id          uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId    uint   `gorm:"school_id" json:"school_id"`                 // 学校ID
	CampusId    uint   `gorm:"campus_id" json:"campus_id"`                 // 学校ID
	StudentName string `gorm:"student_name;size:20"   json:"student_name"` // 学校名称
	Sex         string `gorm:"sex;size:1"   json:"sex"`
	Avatar      string `gorm:"avatar;size:200"   json:"avatar"`
	Address     string `gorm:"address;size:50"   json:"address"`
	SchoolName  string `gorm:"school_name;size:50"   json:"school_name"`
	Birthday    string `gorm:"birthday;size:20" json:"birthday"`
	Remark      string `gorm:"remark;size:100"   json:"remark"`
	Status      string `gorm:"status;size:1"   json:"status"`
	Model
}

func (c *Student) GetStudentList(school_id, campus_id, page uint, limit uint) ([]*Student, int) {
	var student []*Student
	var total int
	db := GetDB()

	// isAdmin := CheckSchoolAdmin(schoolId, user_id)
	db.Table("student").
		Where("school_id = ?", school_id).
		Where("campus_id = ?", campus_id).
		Count(&total).
		Offset((page - 1) * limit).Limit(limit).Find(&student)
	return student, total
}

func (c *Student) UpdateStudent(student_id uint, data map[string]interface{}) bool {
	if err := db.Table("student").Where("id = ?", student_id).Updates(data).Error; err != nil {
		return false
	}
	return true
}

func (c *Student) CreateStudent(student *Student) bool {
	db := GetDB()
	db.Create(&student)
	return true
}
