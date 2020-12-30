/*
 * @Author: 彭林
 * @Date: 2020-12-30 17:12:05
 * @LastEditTime: 2020-12-30 17:28:50
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/service/student_card/type.go
 */
package student_card

type StudentCardParams struct {
	SchoolId   uint `v:"required#学校不能为空"`
	CampusId   uint `v:"required#校区不能为空"`
	CardTypeId uint `v:"required#学员卡id不能为空"`
	StudentId  uint `v:"required#学生不能为空"`
	StartTime  uint
	ExpireTime uint
	Total      uint
	Status     uint `v:"required#状态不能为空"` // 1 未开卡 2 已开卡 3 已作废
}
