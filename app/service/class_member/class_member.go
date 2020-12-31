package class_member

import (
	"errors"
	"goframe-web/app/model"

	"github.com/gogf/gf/os/gtime"
)

func GetClassMemeberList(class_id, status, page, limit uint) (classMemberList []*model.ClassMemberList, total uint, err error) {

	if class_id <= 0 {
		return nil, 0, errors.New("参数错误 class_id")
	}

	if page <= 0 {
		return nil, 0, errors.New("参数错误 page")
	}

	if limit <= 0 {
		return nil, 0, errors.New("参数错误 limit")
	}

	var classMember *model.ClassMember
	classMemberList, total = classMember.GetClassMemberList(class_id, status, page, limit)
	return
}

func CreateClassMember(member *model.ClassMember) (id uint, msg error) {

	if member.SchoolId <= 0 {
		return 0, errors.New("参数错误")
	}

	if member.CampusId <= 0 {
		return 0, errors.New("参数错误")
	}

	if member.ClassId <= 0 {
		return 0, errors.New("参数错误")
	}

	if member.StudentId <= 0 {
		return 0, errors.New("参数错误")
	}

	if member.MemberType == 0 {
		return 0, errors.New("参数错误")
	}

	if member.Status == 0 {
		member.Status = 1
	}

	if member.EntryAt == 0 {
		member.EntryAt = uint(gtime.TimestampMilli())
	}

	var memberModel *model.ClassMember
	memberId := memberModel.IsExistClassMember(member.ClassId, member.StudentId)

	if memberId != 0 {
		return 0, errors.New("学员已存在")
	}

	id = memberModel.CreateClassMember(member)

	return
}

func LeaveClassMember(class_id, student_id, status uint) (id uint, msg error) {

	if class_id <= 0 {
		return 0, errors.New("参数错误")
	}

	if student_id <= 0 {
		return 0, errors.New("参数错误")
	}

	if status <= 0 {
		return 0, errors.New("参数错误")
	}

	var classMember model.ClassMember
	rowId := classMember.IsExistClassMember(class_id, student_id)
	if rowId == 0 {
		return 0, errors.New("学员不存在")
	}
	idDelete := classMember.UpdateClassMember(rowId, status)
	if idDelete {
		return rowId, nil
	}
	return 0, errors.New("删除失败")
}
