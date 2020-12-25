package router

import (
	"goframe-web/app/api/campus"
	"goframe-web/app/api/card"
	"goframe-web/app/api/class"
	"goframe-web/app/api/class_member"
	"goframe-web/app/api/school_user"
	"goframe-web/app/api/student"
	"goframe-web/app/api/subject"
	"goframe-web/app/api/upload_file"
	"goframe-web/app/api/user"
	"goframe-web/app/api/zuowen"
	"goframe-web/app/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.Use(middleware.MiddlewareErrorHandler)
	s.Use(middleware.CORS)
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	s.Group("/api", func(group *ghttp.RouterGroup) {

		group.POST("/login", user.Login)
		group.POST("/signup", user.SignUp)
		group.POST("/zuowen/info", zuowen.GetInfo)

		group.Middleware(middleware.JWTAuth)

		group.GET("/user/info", user.Info)
		group.POST("/signout", user.Signout)
		group.POST("/user/update", user.UpdateInfo)

		group.Group("/school", func(group *ghttp.RouterGroup) {
			group.GET("/list", school_user.GetSchoolList)
			group.GET("/campus/list", campus.GetCampusList)
			group.GET("/campus/simple/list", campus.GetCampusSimpleList)
			group.POST("/campus/update", campus.UpdateCampus)
			group.POST("/campus/create", campus.CreateCampus)
		})

		group.GET("/teacher/list", school_user.GetTeacherList)
		group.POST("/teacher/update", school_user.UpdateTeacher)
		group.POST("/teacher/create", school_user.CreateTeacher)

		group.GET("/class/list", class.GetClassList)
		group.GET("/class/simple/list", class.GetClassSimpleList)
		group.GET("/class/info", class.GetClassInfo)
		group.POST("/class/update", class.UpdateClass)
		group.POST("/class/create", class.CreateClass)
		group.GET("/class/member/list", class_member.GetClassMemeberList)

		group.GET("/student/list", student.GetStudentList)
		group.POST("/student/update", student.UpdateStudent)
		group.POST("/student/create", student.CreateStudent)

		group.GET("/subject/list", subject.GetSubjectList)
		group.POST("/subject/update", subject.UpdateSubject)
		group.POST("/subject/create", subject.CreateSubject)

		group.GET("/card/list", card.GetCardList)
		group.GET("/card/simple/list", card.GetCardSimpleList)
		group.POST("/card/update", card.UpdateCard)
		group.POST("/card/create", card.CreateCard)

		group.POST("/upload/file", upload_file.UploadFile)

		group.POST("/zuowen/sync", zuowen.SaveZuowen)
		group.GET("/zuowen/list", zuowen.GetZuowenList)
		group.GET("/zuowen/lastid", zuowen.GetZuowenLastId)
	})
}
