package model

// Entity is the golang structure for table user.
type Campus struct {
	Id         uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId   uint   `gorm:"school_id" json:"school_id"`               // 学校ID
	CampusName string `gorm:"campus_name;size:20"   json:"campus_name"` // 学校名称
	Address    string `gorm:"address;size:50"   json:"address"`
	Province   string `gorm:"province;size:7"   json:"province"`
	City       string `gorm:"city;size:7"   json:"city"`
	Area       string `gorm:"area;size:7"   json:"area"`
	Model
}

func (c *Campus) GetCampusList(schoolId uint, page uint, Limit uint) ([]*Campus, int) {
	var Campus []*Campus
	var total int
	db := GetDB()
	db.Where("school_id = ?", schoolId).Offset((page - 1) * Limit).Limit(Limit).Find(&Campus)
	db.Table("campus").Where("school_id = ?", schoolId).Count(&total)

	return Campus, total
}

func (c *Campus) GetCampusSimpleList(schoolId uint) ([]*Campus, error) {
	var Campus []*Campus
	db := GetDB()
	db.Where("school_id = ?", schoolId).Find(&Campus)
	return Campus, nil
}

func (c *Campus) UpdateCampus(campusId uint, data map[string]interface{}) bool {
	if err := db.Table("campus").Where("id = ?", campusId).Updates(data).Error; err != nil {
		return false
	}
	return true
}

func (c *Campus) CreateCampus(campus *Campus) bool {
	db := GetDB()
	db.Create(&campus)
	return true
}
