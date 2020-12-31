/*
 * @Author: 彭林
 * @Date: 2020-12-31 15:26:36
 * @LastEditTime: 2020-12-31 15:57:27
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/model/districts.go
 */
package model

type Districts struct {
	Id       uint   `gorm:"AUTO_INCREMENT;size:11" json:"id"`
	Code     uint   `gorm:"code;size:11" json:"code"`
	ParentId uint   `gorm:"parent_id;size:6"   json:"parent_id"`
	Name     string `gorm:"name;type:varchar(32)"   json:"name"`
	FullName string `gorm:"full_name;type:varchar(32)"   json:"full_name"`
	Model
}

func (d *Districts) GetDistrictsList(id uint) []*Districts {
	var dlist []*Districts
	db := GetDB()
	db.Where("parent_id = ?", id).Find(&dlist)
	return dlist
}
