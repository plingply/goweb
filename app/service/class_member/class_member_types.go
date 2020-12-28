/*
 * @Author: 彭林
 * @Date: 2020-12-28 10:57:16
 * @LastEditTime: 2020-12-28 10:59:11
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/service/class_member/class_member_types.go
 */
package class_member

type CreateParam struct {
	SchoolId   uint
	CampusId   uint
	ClassId    uint
	StudentId  uint
	MemberType string
	EntryAt    string
	LeaveAt    string
	Status     uint
}
