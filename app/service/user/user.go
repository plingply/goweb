package user

import (
	"errors"
	"fmt"
	"goframe-web/app/model"
	"goframe-web/library/jwt"
	"goframe-web/library/md5x"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

const (
	SessionMark = "user_info"
)

// 用户注册
func SignUp(param *SignUpParam) (interface{}, error) {
	// 昵称为非必需参数，默认使用账号名称
	if param.Nickname == "" {
		param.Nickname = param.Passport
	}
	// 账号唯一性数据检查
	if ok, _ := CheckPassport(param.Passport); ok {
		return nil, errors.New(fmt.Sprintf("账号 %s 已经存在", param.Passport))
	}

	// 将输入参数赋值到数据库实体对象上
	var user *model.User
	if err := gconv.Struct(param, &user); err != nil {
		return nil, err
	}
	md5str := md5x.GetMD5String(user.Password)
	user.Password = md5str
	result, err := user.Save()

	// 创建usertoken 1
	var userToken model.UserToken
	userToken.UserId = result.Id
	_, e := userToken.Save()

	if e != nil {
		return nil, e
	}

	return result, err
}

// 检查账号是否存在
func CheckPassport(passport string) (bool, uint) {
	var user model.User
	userInfo := user.GetUserInfoByPassport(passport)
	if userInfo.Id != 0 {
		return true, userInfo.Id
	}
	return false, 0
}

// 判断用户是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains(SessionMark)
}

// 更新用户信息
func Update(id uint, reqMap map[string]interface{}) error {
	lu, err := GetUserInfo(id)
	if err != nil {
		return err
	}
	reqMap["password"] = md5x.GetMD5String(reqMap["password"].(string))
	_, err = lu.Update(reqMap)
	return err
}

// 获得用户信息详情
func GetUserInfo(id uint) (*model.UserRoles, error) {
	var user model.UserRoles
	if id == 0 {
		return nil, errors.New("参数错误")
	}
	result, err := user.GetUserInfoById(id)
	return result, err
}

// 登录系统
func Login(passport string, password string) (token string, err error) {

	if passport == "" || password == "" {
		return "", errors.New("参数错误")
	}

	var user model.User
	userInfo := user.GetUserInfoByPassport(passport)
	if userInfo.Id == 0 {
		return "", errors.New("用户不存在")
	}
	password = md5x.GetMD5String(password)
	if userInfo.Password != password {
		return "", errors.New("账号或密码错误")
	}

	var claims jwt.CustomClaims
	claims.Username = userInfo.Passport
	claims.Userid = userInfo.Id
	claims.StandardClaims = jwtgo.StandardClaims{
		NotBefore: int64(time.Now().Unix() - 1000),        // 签名生效时间
		ExpiresAt: int64(time.Now().Unix() + 3600*24*365), // 过期时间 一小时
		Issuer:    jwt.GetSignKey(),                       //签名的发行者
	}
	mjwt := jwt.NewJWT()
	token, err = mjwt.CreateToken(claims)

	// 创建usertoken 1
	var userToken *model.UserToken
	userToken = userToken.GetUserTokenByUserId(userInfo.Id)

	if userToken.Id == 0 {
		// 创建usertoken 1
		userToken.UserId = userInfo.Id
		userToken.Save()
	}

	// 更新 token
	_, e := userToken.Update(userInfo.Id, token)
	if e != nil {
		err = e
		return
	}
	return
}

// 退出
func Signout(userid uint) error {
	// 更新 token
	var usertoken model.UserToken
	_, e := usertoken.Update(userid, "")
	if e != nil {
		return e
	}
	return nil
}
