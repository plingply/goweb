/*
 * @Author: 彭林
 * @Date: 2020-12-24 15:03:31
 * @LastEditTime: 2020-12-25 13:25:13
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/model/zuowen.go
 */
package model

type Zuowen struct {
	Id         uint   `gorm:"AUTO_INCREMENT" json:"id"`
	ZuowenId   uint   `gorm:"zuowen_id" json:"zuowen_id"`
	Comments   string `gorm:"comments;size:5000"   json:"comments"`
	Genre      string `gorm:"genre;size:20"   json:"genre"`
	Grade      string `gorm:"grade;size:20"   json:"grade"`
	Quality    string `gorm:"quality;size:10"   json:"quality"`
	ReadCount  uint   `gorm:"read_count"   json:"read_count"`
	Theme      string `gorm:"theme;size: 100"   json:"theme"`
	Title      string `gorm:"title;size: 100"   json:"title"`
	WordNumber uint   `gorm:"word_number"   json:"word_number"`
	ZwContent  string `gorm:"zw_content;type:text(60000)"   json:"zw_content"`
	Model
}

// gener string, grade string,
func (z *Zuowen) GetZuowenList(page, limit uint) ([]*Zuowen, int) {
	var zuowen []*Zuowen
	var total int
	db := GetDB()
	db.Table("zuowen").Order("zuowen_id desc").Count(&total).Offset((page - 1) * limit).Limit(limit).Find(&zuowen)
	return zuowen, total
}
