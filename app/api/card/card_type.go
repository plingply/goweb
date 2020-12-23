package card

type CardRequest struct {
	CardName string `v:"required|length:1,10#学员卡名称不能为空|名称长度应当在:min到:max之间"`
	SchoolID uint   `v:"required#学校不能为空"`
	CampusID uint   `v:"required#校区不能为空"`
	Remark   string `v:"length:0,100#备注长度应当在:min到:max之间"`
	Status   string
}
