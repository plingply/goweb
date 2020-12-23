package class_member

import (
	"errors"
	"goframe-web/app/model"
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
