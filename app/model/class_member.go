package model

type ClassMember struct {
	Id         uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId   uint   `gorm:"school_id" json:"school_id"` // 学校ID
	CampusId   uint   `gorm:"campus_id" json:"campus_id"` // 学校ID
	ClassId    uint   `gorm:"class_id" json:"class_id"`
	StudentId  uint   `gorm:"student_id" json:"student_id"`
	MemberType string `gorm:"member_type;size:1"   json:"member_type"`
	EntryAt    string `gorm:"entry_at;size:20" json:"entry_at"`
	LeaveAt    string `gorm:"leave_at;size:20" json:"leave_at"`
	Status     uint   `gorm:"status;size:1"   json:"status"`
	Model
}

type ClassMemberList struct {
	Id          uint   `json:"id"`
	SchoolId    uint   `json:"school_id"` // 学校ID
	CampusId    uint   `json:"campus_id"` // 学校ID
	ClassId     uint   `json:"class_id"`
	StudentId   uint   `json:"student_id"`
	StudentName string `json:"student_name"`
	MemberType  string `json:"member_type"`
	EntryAt     string `json:"entry_at"`
	LeaveAt     string `json:"leave_at"`
	Status      uint   `json:"status"`
	Model
}

func (c *ClassMember) GetClassMemberList(school_id, campus_id, class_id, page, limit uint) (classMemberList []*ClassMemberList, total uint) {
	db := GetDB()
	// isAdmin := CheckSchoolAdmin(schoolId, user_id)
	db.Table("class_member").
		Select("class_member.*, student.student_name").
		Where("class_member.school_id = ?", school_id).
		Where("class_member.campus_id = ?", campus_id).
		Where("class_member.class_id = ?", class_id).
		Joins("left join student on class_member.student_id = student.id").
		Count(&total).
		Offset((page - 1) * limit).Limit(limit).Scan(&classMemberList)
	return classMemberList, total
}

func (c *ClassMember) CreateClassMember(member *ClassMember) (id uint) {
	db := GetDB()
	db.Create(&member)
	return member.Id
}

func (c *ClassMember) IsExistClassMember(class_id, student_id uint) bool {
	db := GetDB()
	var member ClassMember
	db.Where("class_id = ?", class_id).Where("student_id = ?", student_id).Find(&member)
	if member.Id == 0 {
		return false
	}
	return true
}
