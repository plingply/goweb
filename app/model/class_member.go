package model

import "github.com/gogf/gf/os/gtime"

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

func (c *ClassMember) GetClassMemberList(class_id, status, page, limit uint) (classMemberList []*ClassMemberList, total uint) {
	db := GetDB()
	// isAdmin := CheckSchoolAdmin(schoolId, user_id)
	db = db.Table("class_member").
		Select("class_member.*, student.student_name").
		Where("class_member.class_id = ?", class_id)
	if status != 0 {
		db = db.Where("class_member.status = ?", status)
	}
	db = db.Where("class_member.deleted_at IS NULL").
		Joins("left join student on class_member.student_id = student.id").
		Count(&total)
	if page > 0 && limit > 0 {
		db = db.Offset((page - 1) * limit).Limit(limit)
	}
	db.Scan(&classMemberList)
	return classMemberList, total
}

func (c *ClassMember) CreateClassMember(member *ClassMember) (id uint) {
	db := GetDB()
	db.Create(&member)
	return member.Id
}

func (c *ClassMember) UpdateClassMember(id, status uint) bool {
	db := GetDB()
	var member ClassMember
	db = db.Model(&member).Where("id = ?", id)
	if status == 2 {
		db = db.Update("leave_at", gtime.TimestampMilliStr())
	}
	db.Update("status", status)
	return true
}

func (c *ClassMember) IsExistClassMember(class_id, student_id uint) uint {
	db := GetDB()
	var member ClassMember
	db.Where("class_id = ?", class_id).Where("student_id = ?", student_id).Find(&member)
	return member.Id
}
