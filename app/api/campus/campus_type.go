package campus

// 注册请求参数，用于前后端交互参数格式约定
type CampusRequest struct {
	CampusName string `v:"required|length:3,20#校区名称不能为空|名称长度应当在:min到:max之间"`
	Address    string `v:"required|length:1,50#地址不能为空|名称长度应当在:min到:max之间"`
	Province   string `v:"required|length:0,7#请选择省|省份长度超限"`
	City       string `v:"required|length:0,7#请选择市|市长度超限"`
	Area       string `v:"required|length:0,7#请选择区|区长度超限"`
}
