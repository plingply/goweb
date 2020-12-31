package model

import (
	"fmt"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Logger struct {
}

func GetDB() *gorm.DB {
	return db
}

func (logger Logger) Print(values ...interface{}) {
	g.Log("sqlogger").Error(values)
}

func init() {
	initDBtoDefault()
}

func initDBtoDefault() {
	var err error
	user := g.Cfg().GetString("database.user")
	databasetype := g.Cfg().GetString("database.type")
	password := g.Cfg().GetString("database.pwd")
	host := g.Cfg().GetString("database.host")
	port := g.Cfg().GetString("database.port")
	tablename := g.Cfg().GetString("database.tablename")
	urlStr := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + tablename + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open(databasetype, urlStr)

	if err != nil {
		fmt.Println(err)
		panic("数据库连接失败")
	}

	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(1 * time.Second)
	db.DB().SetMaxIdleConns(20)   //最大打开的连接数
	db.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	db.SingularTable(true)        //表生成结尾不带s
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// db.SetLogger(Logger{})
	Createtable()
}

// Createtable 初始化表 如果不存在该表 则自动创建
func Createtable() {
	GetDB().AutoMigrate(
		&User{},
		&UserToken{},
		&School{},
		&Campus{},
		&SchoolUser{},
		&Classs{},
		&Student{},
		&Subject{},
		&Card{},
		&ClassMember{},
		&Zuowen{},
		&Course{},
		&Peotry{},
		&NoteList{},
		&CourseMember{},
		&StudentCard{},
		&Districts{},
	)
}

type Model struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
