package model

// Entity is the golang structure for table user.
type SchoolUser struct {
	Id          uint   `gorm:"AUTO_INCREMENT" json:"id"` // 学校ID
	UserId      uint   `gorm:"user_id" json:"user_id"`
	CampusId    uint   `gorm:"campus_id" json:"campus_id"`
	SchoolId    uint   `gorm:"school_id" json:"school_id"` // 学校ID
	TeacherName string `gorm:"teacher_name;size:10" json:"teacher_name"`
	Sex         string `gorm:"sex;size:1" json:"sex"`
	Phone       string `gorm:"phone;size:11" json:"phone"`
	Address     string `gorm:"address;size:50" json:"address"`
	Identity    string `gorm:"identity;size:10" json:"identity"` // school, campus
	Birthday    string `gorm:"birthday;size:20" json:"birthday"`
	EntryAt     string `gorm:"entry_at;size:20" json:"entry_at"`
	Model
}

func (c *SchoolUser) GetTeacherList(schoolId, campusId, page, limit uint) ([]*SchoolUser, int) {
	var schoolUser []*SchoolUser
	var total int
	db := GetDB()
	db = db.Table("school_user").Where("school_id = ?", schoolId).Count(&total)
	if campusId != 0 {
		db = db.Where("campus_id = ?", campusId)
	}
	if page > 0 && limit > 0 {
		db = db.Offset((page - 1) * limit).Limit(limit)
	}
	db.Find(&schoolUser)

	return schoolUser, total
}

func (c *SchoolUser) UpdateTeacher(teacherId uint, data map[string]interface{}) bool {
	db.Table("school_user").Where("id = ?", teacherId).Updates(data)
	return true
}

func (c *SchoolUser) CreateTeacher(schoolUser *SchoolUser) bool {
	db := GetDB()
	db.Create(&schoolUser)
	return true
}

func (c *SchoolUser) GetTeacherInfoByPhone(campus_id uint, phone string) *SchoolUser {
	var user SchoolUser
	db := GetDB()
	db.Where("phone = ?", phone).Where("campus_id = ?", campus_id).First(&user)
	return &user
}

// 检查老师账号是否存在
func (c *SchoolUser) CheckTeacher(campus_id uint, phone string) bool {
	var user SchoolUser
	teacher := user.GetTeacherInfoByPhone(campus_id, phone)
	if teacher.Id != 0 {
		return true
	}
	return false
}
