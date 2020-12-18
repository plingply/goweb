package model

// Entity is the golang structure for table user.
type CampusSchool struct {
	Id       uint `gorm:"AUTO_INCREMENT" json:"id"`
	CampusId uint `gorm:"campus_id" json:"campus_id"`
	SchoolId uint `gorm:"school_id" json:"school_id"` // 学校ID
	Model
}
