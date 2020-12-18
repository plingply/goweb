package model

// Entity is the golang structure for table user.
type SchoolUser struct {
	Id          uint   `gorm:"AUTO_INCREMENT" json:"id"` // 学校ID
	UserId      uint   `gorm:"user_id" json:"user_id"`
	CampusId    uint   `gorm:"campus_id" json:"campus_id"`
	SchoolId    uint   `gorm:"school_id" json:"school_id"` // 学校ID
	TeacherName string `gorm:"teacher_name" json:"teacher_name"`
	Model
}
