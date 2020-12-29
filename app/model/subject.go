package model

type Subject struct {
	Id          uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId    uint   `gorm:"school_id" json:"school_id"`                 // 学校ID
	CampusId    uint   `gorm:"campus_id" json:"campus_id"`                 // 学校ID
	SubjectName string `gorm:"subject_name;size:20"   json:"subject_name"` // 学校名称
	Remark      string `gorm:"remark;size:100"   json:"remark"`            // 学校名称
	Status      string `gorm:"status;size:1"   json:"status"`
	Model
}

func (c *Subject) GetSubjectList(school_id, campus_id, page, limit uint) ([]*Subject, int) {
	var subject []*Subject
	var total int
	db := GetDB()

	// isAdmin := CheckSchoolAdmin(schoolId, user_id)
	db = db.Table("subject").
		Where("school_id = ?", school_id).
		Where("campus_id = ?", campus_id).
		Count(&total)
	if page > 0 && limit > 0 {
		db = db.Offset((page - 1) * limit).Limit(limit)
	}
	db.Find(&subject)
	return subject, total
}

func (c *Subject) UpdateSubject(subject_id uint, data map[string]interface{}) bool {
	if err := db.Table("subject").Where("id = ?", subject_id).Updates(data).Error; err != nil {
		return false
	}
	return true
}

func (c *Subject) CreateSubject(subject *Subject) bool {
	db := GetDB()
	db.Create(&subject)
	return true
}
