package model

type Classs struct {
	Id        uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId  uint   `gorm:"school_id" json:"school_id"`             // 学校ID
	CampusId  uint   `gorm:"campus_id" json:"campus_id"`             // 学校ID
	ClassName string `gorm:"class_name;size:20"   json:"class_name"` // 学校名称
	ClassType string `gorm:"class_type;size:1"   json:"class_type"`
	Status    string `gorm:"status;size:1"   json:"status"`
	Model
}

func (c *Classs) GetClassList(school_id, campus_id, user_id uint, page uint, limit uint) ([]*Classs, int) {
	var class []*Classs
	var total int
	db := GetDB()

	// isAdmin := CheckSchoolAdmin(schoolId, user_id)
	db.Where("school_id = ?", school_id).Where("campus_id = ?", campus_id).Offset((page - 1) * limit).Limit(limit).Find(&class)
	db.Table("classs").Where("school_id = ?", school_id).Where("campus_id = ?", campus_id).Count(&total)

	return class, total
}

func (c *Classs) GetClassSimpleList(school_id, campus_id, user_id uint) []*Classs {
	var class []*Classs
	db := GetDB()

	// isAdmin := CheckSchoolAdmin(schoolId, user_id)
	db.Where("school_id = ?", school_id).Where("campus_id = ?", campus_id).Find(&class)
	return class
}

func (c *Classs) UpdateClass(class_id uint, data map[string]interface{}) bool {
	if err := db.Table("classs").Where("id = ?", class_id).Updates(data).Error; err != nil {
		return false
	}
	return true
}

func (c *Classs) CreateClass(class *Classs) bool {
	db := GetDB()
	db.Create(&class)
	return true
}
