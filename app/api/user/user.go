package user

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"goframe-web/app/service/user"
	"goframe-web/library/response"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

func SignUp(r *ghttp.Request) {
	var (
		data        *SignUpRequest
		signUpParam *user.SignUpParam
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &signUpParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if result, err := user.SignUp(signUpParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}

// 用户详情
func Info(r *ghttp.Request) {
	id := r.GetQueryUint("id")

	if id == 0 {
		id = r.GetCtxVar("user_id").Uint()
	}

	if result, err := user.GetUserInfo(id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}

// 修改用户信息
func UpdateInfo(r *ghttp.Request) {

	var reqMap = make(map[string]interface{})
	id := r.GetFormUint("id")

	if r.GetFormBool("nickname") {
		reqMap["nickname"] = r.GetFormString("nickname")
	}

	if r.GetFormBool("password") {
		reqMap["password"] = r.GetFormString("password")
	}

	if err := user.Update(id, reqMap); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", "修改成功")
	}
}

// 登录
func Login(r *ghttp.Request) {
	passport := r.GetFormString("passport")
	password := r.GetFormString("password")
	if result, err := user.Login(passport, password); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}

// 退出
func Signout(r *ghttp.Request) {
	userid := r.GetCtxVar("user_id").Uint()
	if err := user.Signout(userid); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", nil)
	}
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func objKeySort(data map[string]string) ([]string, []string) {

	var t []string
	var r []string
	for k, _ := range data {
		t = append(t, k)
	}

	sort.Strings(t)

	for _, v := range t {
		r = append(r, data[v])
	}
	return t, r
}

func getSign(data map[string]string) string {
	t := md5V(data["date"])
	_, v := objKeySort(data)
	fmt.Println("ttt:", t, len(t), v)
	r := ""
	for _, val := range v {
		r += val
	}
	r += "CVrFcVJu564ArjsShh058xSjb"
	return t[0:16] + md5V(r) + t[16:32]
}

func GetArt(r *ghttp.Request) {

	ciphertext := make(map[string]string)
	ciphertext["date"] = gtime.Datetime()
	ciphertext["zuoWenId"] = "1"

	urlValues := url.Values{}
	urlValues.Add("zuoWenId", "1")
	urlValues.Add("date", gtime.Datetime())
	urlValues.Add("ciphertext", getSign(ciphertext))

	fmt.Println("urlValues:", urlValues)

	gtime.SetTimeZone("Asia/Shanghai")
	var postUrl = "https://api.xhzapp.com/api/v5/AppBook/GetZuoWenModel?r=Thu Dec 24 2020 13:23:54 GMT 0800 (中国标准时间)"

	fmt.Println("postUrl:", postUrl)
	res, err := http.PostForm(postUrl, urlValues)
	defer res.Body.Close()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	var result map[string]interface{}
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	fmt.Println("result:", result)
	if result["code"].(string) == "0" {
		response.JsonExit(r, 1, result["msg"].(string))
	}
	response.JsonExit(r, 0, "ok", result["data"])

}
