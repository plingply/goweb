/*
 * @Author: 彭林
 * @Date: 2020-12-30 17:05:32
 * @LastEditTime: 2020-12-30 17:58:01
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/service/student_card/student_card.go
 */
package student_card

import (
	"errors"
	"goframe-web/app/model"

	"github.com/gogf/gf/os/gtime"
)

func ActivateCard(params *StudentCardParams) (uint, error) {

	var modelStudentCard model.StudentCard

	isExist := modelStudentCard.IsExistStudentCard(params.StudentId, params.CardTypeId)

	if isExist {
		return 0, errors.New("改学员已存在相同类型的学员卡")
	}
	modelStudentCard.SchoolId = params.SchoolId
	modelStudentCard.CampusId = params.CampusId
	modelStudentCard.CardTypeId = params.CardTypeId
	modelStudentCard.StudentId = params.StudentId
	modelStudentCard.ExpireTime = params.ExpireTime
	modelStudentCard.Total = params.Total
	modelStudentCard.Residue = params.Total
	modelStudentCard.Status = params.Status
	if params.Status == 2 {
		modelStudentCard.StartTime = uint(gtime.Millisecond())
	}

	err := modelStudentCard.CreateStudentCard(&modelStudentCard)
	if err != nil {
		return 0, err
	}
	return modelStudentCard.Id, nil
}
