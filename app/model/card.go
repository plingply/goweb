package model

type Card struct {
	Id       uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId uint   `gorm:"school_id" json:"school_id"`           // 学校ID
	CampusId uint   `gorm:"campus_id" json:"campus_id"`           // 学校ID
	CardName string `gorm:"card_name;size:20"   json:"card_name"` // 学校名称
	Remark   string `gorm:"remark;size:100"   json:"remark"`
	Status   string `gorm:"status;size:1"   json:"status"`
	Model
}

func (c *Card) GetCardList(school_id, campus_id, page uint, limit uint) ([]*Card, int) {
	var card []*Card
	var total int
	db := GetDB()

	// isAdmin := CheckSchoolAdmin(schoolId, user_id)
	db.Table("card").Where("school_id = ?", school_id).Where("campus_id = ?", campus_id).Count(&total).Offset((page - 1) * limit).Limit(limit).Find(&card)
	return card, total
}

func (c *Card) GetCardSimpleList(school_id, campus_id uint) []*Card {
	var card []*Card
	db := GetDB()

	// isAdmin := CheckSchoolAdmin(schoolId, user_id)
	db.Where("school_id = ?", school_id).Where("campus_id = ?", campus_id).Find(&card)
	return card
}

func (c *Card) UpdateCard(card_id uint, data map[string]interface{}) bool {
	if err := db.Table("card").Where("id = ?", card_id).Updates(data).Error; err != nil {
		return false
	}
	return true
}

func (c *Card) CreateCard(card *Card) bool {
	db := GetDB()
	db.Create(&card)
	return true
}
