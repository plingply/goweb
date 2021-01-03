package school_user

// 注册请求参数，用于前后端交互参数格式约定
type TeacherRequest struct {
	TeacherName string `v:"required|length:1,20#老师名称不能为空|名称长度应当在:min到:max之间"`
	SchoolId    uint   `v:"required#学校不能为空"`
	CampusId    uint
	Sex         string `v:"required|length:1,1#请选择性别|性别非法"`
	Phone       string `v:"required|length:11,11#联系电话不能为空|联系电话格式错误"`
	Identity    string `v:"required#身份类型错误"`
	Address     string
	Birthday    string
	EntryAt     string
}
