package class_member

import (
	"errors"
	"goframe-web/app/model"

	"github.com/gogf/gf/os/gtime"
)

func GetClassMemeberList(school_id, campus_id, class_id, page, limit uint) (classMemberList []*model.ClassMemberList, total uint, err error) {

	if school_id <= 0 {
		return nil, 0, errors.New("参数错误 school_id")
	}

	if campus_id <= 0 {
		return nil, 0, errors.New("参数错误 campus_id")
	}

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
	classMemberList, total = classMember.GetClassMemberList(school_id, campus_id, class_id, page, limit)
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

	if member.MemberType == "" {
		return 0, errors.New("参数错误")
	}

	if member.Status == 0 {
		member.Status = 1
	}

	if member.EntryAt == "" {
		member.EntryAt = gtime.TimestampMilliStr()
	}

	var memberModel *model.ClassMember
	isExist := memberModel.IsExistClassMember(member.ClassId, member.StudentId)

	if isExist {
		return 0, errors.New("学员已存在")
	}

	id = memberModel.CreateClassMember(member)

	return
}
