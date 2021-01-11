package model

import (
	"fmt"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var db *gorm.DB

type Logger struct {
}

func GetDB() *gorm.DB {
	return db
}

func (logger Logger) Print(v ...interface{}) {
	g.Log("sqlogger").Info("sql-log:", v)
	switch v[0] {
	case "sql":
		var lmap = []interface{}{
			"sql",
			zap.String("module", "gorm"),
			zap.String("type", "sql"),
			zap.Any("src", v[1]),
			zap.Any("duration", v[2]),
			zap.Any("sql", v[3]),
			zap.Any("values", v[4]),
			zap.Any("rows_returned", v[5]),
		}
		g.Log("sqlogger").Print(lmap)
	case "log":
		g.Log("sqlogger").Info(zap.Any("gorm", v[2]))
	}
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
	db.SetLogger(Logger{})
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
