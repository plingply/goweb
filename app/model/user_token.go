package model

type UserToken struct {
	Id       uint        `gorm:"AUTO_INCREMENT" json:"id"`        // 用户ID
	UserId   uint        `json:"user_id"`        // 用户ID
	Token    string      `gorm:"default:''" json:"token"`  // token
	Model
}


func (u *UserToken)Save() (*UserToken, error)  {
	db := GetDB()
	db.Create(&u)
	return u, nil
}

func (u *UserToken)Update(userid uint, token string)(*UserToken, error){
	db := GetDB()
	db.Model(u).Where("user_id = ?", userid).Update("token", token)
	return u, nil
}

func (u *UserToken)GetUserTokenByUserId(userid uint) *UserToken  {
	var user UserToken
	db := GetDB()
	db.Where("user_id = ?", userid).First(&user)
	return &user
}

func (u *UserToken)Vtoken(userId uint, token string) bool{
	ustoken := u.GetUserTokenByUserId(userId)

	if ustoken.Token != token {
		return false
	}

	return true
}