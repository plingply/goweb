package router

import (
	"goframe-web/app/api/campus"
	"goframe-web/app/api/card"
	"goframe-web/app/api/class"
	"goframe-web/app/api/class_member"
	"goframe-web/app/api/course"
	"goframe-web/app/api/districts"
	"goframe-web/app/api/peotry"
	"goframe-web/app/api/school"
	"goframe-web/app/api/school_user"
	"goframe-web/app/api/student"
	"goframe-web/app/api/student_card"
	"goframe-web/app/api/subject"
	"goframe-web/app/api/upload_file"
	"goframe-web/app/api/user"
	"goframe-web/app/api/wechat"
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
		group.POST("/peotry/info", peotry.GetInfo)
		group.GET("/city/:parent_id", districts.GetDistrictsList)
		group.GET("/wechat/verification", wechat.Verification)

		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.GET("/info", user.Info)
			group.POST("/signout", user.Signout)
			group.POST("/update", user.UpdateInfo)
		})

		group.Group("/school", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.GET("/list", school_user.GetSchoolList)
			group.POST("/create", school.CreateSchool)
			group.GET("/campus/list", campus.GetCampusList)
			group.GET("/campus/simple/list", campus.GetCampusSimpleList)
			group.POST("/campus/update", campus.UpdateCampus)
			group.POST("/campus/create", campus.CreateCampus)
		})

		s.Group("/teacher", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.GET("/list", school_user.GetTeacherList)
			group.POST("/update", school_user.UpdateTeacher)
			group.POST("/create", school_user.CreateTeacher)
		})

		s.Group("/class", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.GET("/list", class.GetClassList)
			group.GET("/simple/list", class.GetClassSimpleList)
			group.GET("/info", class.GetClassInfo)
			group.POST("/update", class.UpdateClass)
			group.POST("/create", class.CreateClass)
			group.GET("/member/list", class_member.GetClassMemeberList)
			group.POST("/member/create", class_member.CreateClassMemeber)
			group.POST("/member/leave", class_member.LeaveClassMember)
		})

		s.Group("/student", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.GET("/list", student.GetStudentList)
			group.POST("/update", student.UpdateStudent)
			group.POST("/create", student.CreateStudent)
			group.POST("/student/card/activate", student_card.ActivateCard)
		})

		s.Group("/subject", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.GET("/list", subject.GetSubjectList)
			group.POST("/update", subject.UpdateSubject)
			group.POST("/create", subject.CreateSubject)
		})

		s.Group("/card", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.GET("/list", card.GetCardList)
			group.GET("/simple/list", card.GetCardSimpleList)
			group.POST("/update", card.UpdateCard)
			group.POST("/create", card.CreateCard)
		})

		s.Group("/upload", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.POST("/file", upload_file.UploadFile)
		})

		s.Group("/zuowen", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.POST("/sync", zuowen.SaveZuowen)
			group.GET("/list", zuowen.GetZuowenList)
			group.GET("/lastid", zuowen.GetZuowenLastId)
		})

		s.Group("/peotry", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.POST("/sync", peotry.CreatePeotry)
			group.GET("/lastid", peotry.GetPeotryLastId)
			group.GET("/list", peotry.GetPeotryList)
		})

		s.Group("/course", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.JWTAuth)
			group.GET("/list", course.GetCourseList)
			group.POST("/check", course.CheckCourse)
			group.POST("/add", course.AddCourse)
		})

	})
}
