package class_member

type createClassMemberRequest struct {
	SchoolId   uint   `v:"required|min:1#学校id不能为空|非法参数SchoolId"`
	CampusId   uint   `v:"required|min:1#校区id不能为空|非法参数CampusId"`
	ClassId    uint   `v:"required|min:1#班级id不能为空|非法参数ClassId"`
	StudentId  uint   `v:"required|min:1#学员id不能为空|非法参数StudentId"`
	MemberType string `v:"required#学生类型不能为空"`
	EntryAt    string
	LeaveAt    string
	Status     uint
}
