package student

type StudentRequest struct {
	StudentName string `v:"required|length:1,10#学生姓名不能为空|名称长度应当在:min到:max之间"`
	SchoolID    uint   `v:"required#学校不能为空"`
	CampusID    uint   `v:"required#校区不能为空"`
	Sex         string
	Avatar      string
	Address     string
	SchoolName  string
	Birthday    string
	Remark      string
	Status      string
}
