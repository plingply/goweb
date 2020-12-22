package class

type ClassRequest struct {
	ClassName string `v:"required|length:1,20#班级名称不能为空|名称长度应当在:min到:max之间"`
	SchoolId  uint   `v:"required#学校不能为空"`
	CampusId  uint   `v:"required#校区不能为空"`
	ClassType string `v:"required#班级类型不能为空"`
	Status    string `v:"required#班级状态不能为空"`
}
