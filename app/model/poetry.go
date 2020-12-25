/*
 * @Author: 彭林
 * @Date: 2020-12-25 14:17:21
 * @LastEditTime: 2020-12-25 15:46:11
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/model/poetry.go
 */
package model

type Peotry struct {
	Id          uint   `gorm:"AUTO_INCREMENT" json:"id"`
	PeotryId    uint   `gorm:"peotry_id" json:"peotry_id"`
	Align       uint   `gorm:"align"   json:"align"`
	Appreciate  string `gorm:"appreciate;type:text(60000)"   json:"appreciate"`
	Author      string `gorm:"author;size:20"   json:"author"`
	AuthorMore  string `gorm:"author_more;size:2000"   json:"author_more"`
	DuYin       string `gorm:"du_yin;size:512"   json:"du_yin"`
	OrgText     string `gorm:"org_text;type:text(10000)"   json:"org_text"`
	Reason      string `gorm:"reason;type:text(10000)"   json:"reason"`
	Title       string `gorm:"title;size: 100"   json:"title"`
	Translation string `gorm:"translation;type:text(60000)"   json:"translation"`
	Video       string `gorm:"video;size:512"   json:"video"`
	Years       string `gorm:"years;size:20"   json:"years"`
	Model
}

type NoteList struct {
	Id       uint   `gorm:"AUTO_INCREMENT" json:"id"`
	PeotryId uint   `gorm:"peotry_id" json:"peotry_id"`
	Ciyu     string `gorm:"ciyu;size:20"   json:"ciyu"`
	Notes    string `gorm:"notes;size:2000"   json:"notes"`
	Model
}

type PeotryParams struct {
	Peotry
	NoteList []*NoteList
}

func (c *Peotry) CreatePeotry(peotry *Peotry) uint {
	db := GetDB()
	db.Create(&peotry)
	return peotry.Id
}

func (c *NoteList) CreateNoteList(notelist *NoteList) uint {
	db := GetDB()
	db.Create(&notelist)
	return notelist.Id
}
