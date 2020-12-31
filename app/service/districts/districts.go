/*
 * @Author: 彭林
 * @Date: 2020-12-31 15:33:53
 * @LastEditTime: 2020-12-31 15:57:06
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/service/districts/districts.go
 */
package districts

import (
	"errors"
	"goframe-web/app/model"
)

func GetDistrictsList(id uint) ([]*model.Districts, error) {
	if id == 0 {
		return nil, errors.New("参数错误")
	}
	var districts model.Districts
	dlist := districts.GetDistrictsList(id)
	return dlist, nil
}
