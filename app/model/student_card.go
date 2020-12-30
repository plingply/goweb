/*
 * @Author: 彭林
 * @Date: 2020-12-30 14:46:05
 * @LastEditTime: 2020-12-30 15:07:35
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/model/student_card.go
 */
package model

type StudentCard struct {
	Id         uint   `gorm:"AUTO_INCREMENT" json:"id"`
	SchoolId   uint   `gorm:"school_id" json:"school_id"`
	CampusId   uint   `gorm:"campus_id" json:"campus_id"`
	CardTypeId uint   `gorm:"card_type_id" json:"card_type_id"`
	StudentId  uint   `gorm:"student_id" json:"student_id"`
	StartTime  uint   `gorm:"start_time;default:0;type:BIGINT"   json:"start_time"`
	ExpireTime uint   `gorm:"expire_time;default:0;type:BIGINT"   json:"expire_time"`
	Total      uint   `gorm:"total;default:0;type:BIGINT"   json:"total"`
	Residue    uint   `gorm:"residue;default:0;type:BIGINT"   json:"residue"`
	Status     string `gorm:"status;size:1"   json:"status"` // 1 未开卡 2 已开卡 3 已作废
	Model
}
