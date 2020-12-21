package model

// Entity is the golang structure for table user.
type School struct {
	Id         uint   `gorm:"AUTO_INCREMENT" json:"id"`                 // 学校ID
	SchoolName string `gorm:"school_name;size:20"   json:"school_name"` // 学校名称
	Logo       string `gorm:"logo;size:255"   json:"logo"`              // 用户密码
	Model
}
