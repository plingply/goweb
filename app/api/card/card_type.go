package card

type CardRequest struct {
	CardName string `v:"required|length:1,10#学员卡名称不能为空|名称长度应当在:min到:max之间"`
	SchoolId uint   `v:"required#学校不能为空"`
	CampusId uint   `v:"required#校区不能为空"`
	Remark   string `v:"length:0,100#备注长度应当在:min到:max之间"`
	Status   uint
	CardType uint `v:"required#学员卡类型不能为空"`
}
