/*
 * @Author: 彭林
 * @Date: 2020-12-25 14:39:43
 * @LastEditTime: 2020-12-25 15:05:09
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/service/peotry/peotry.go
 */
package peotry

import "goframe-web/app/model"

func CreatePeotry(peotry *model.Peotry, note_list []*model.NoteList) uint {

	db := model.GetDB()

	var oldpeotry model.Peotry
	db.Where("peotry_id = ?", peotry.PeotryId).Find(&oldpeotry)

	if oldpeotry.Id != 0 {
		return 0
	}

	peotry.CreatePeotry(peotry)
	CreateNoteList(note_list)
	return peotry.PeotryId
}

func CreateNoteList(note_list []*model.NoteList) {
	for _, v := range note_list {
		v.CreateNoteList(v)
	}
}

func GetPeotryLastId() int {

	db := model.GetDB()
	var peotrys []model.Peotry
	db.Order("peotry_id desc").Limit(1).Find(&peotrys)

	if len(peotrys) == 0 {
		return 0
	}
	return int(peotrys[0].PeotryId)
}
