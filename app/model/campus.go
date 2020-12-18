package model

// Entity is the golang structure for table user.
type Campus struct {
	Id         uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId   uint   `gorm:"school_id" json:"school_id"`       // 学校ID
	CampusName string `gorm:"campus_name"   json:"campus_name"` // 学校名称
	Address    string `gorm:"adress"   json:"adress"`
	Province   string `gorm:"province"   json:"province"`
	City       string `gorm:"city"   json:"city"`
	Area       string `gorm:"area"   json:"area"`
	Model
}
