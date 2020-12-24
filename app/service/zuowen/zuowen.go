/*
 * @Author: 彭林
 * @Date: 2020-12-24 14:59:06
 * @LastEditTime: 2020-12-24 19:30:51
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/service/zuowen/zuowen.go
 */
package zuowen

import (
	"errors"
	"goframe-web/app/model"
)

func CreateZuowen(zuowen *model.Zuowen) uint {

	db := model.GetDB()

	var zuowens model.Zuowen
	db.Where("zuowen_id = ?", zuowen.ZuowenId).Find(&zuowens)

	if zuowens.Id != 0 {
		return 0
	}
	db.Create(&zuowen)
	return zuowen.ZuowenId
}

func GetZuowenList(page, limit uint) (result interface{}, total int, err error) {

	if page <= 0 {
		return nil, 0, errors.New("参数错误")
	}
	if limit <= 0 {
		return nil, 0, errors.New("参数错误")
	}

	var zuowen *model.Zuowen
	result, total = zuowen.GetZuowenList(page, limit)

	return result, total, nil
}

func GetZuowenLastId() (id int, err error) {
	db := model.GetDB()
	var zuowens []model.Zuowen
	db.Order("zuowen_id desc").Limit(1).Find(&zuowens)

	if len(zuowens) == 0 {
		return 0, nil
	}
	return int(zuowens[0].ZuowenId), nil
}

func GetInfo(zuowen_id uint) (*model.Zuowen, error) {

	if zuowen_id <= 0 {
		return nil, errors.New("参数错误")
	}

	db := model.GetDB()
	var zuowen model.Zuowen
	db.Where("zuowen_id = ?", zuowen_id).Find(&zuowen)

	return &zuowen, nil
}
