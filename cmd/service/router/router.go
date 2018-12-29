package router

import (
	"github.com/gin-gonic/gin"
	"ojbk.io/gopherCron/cmd/service/controller/etcd_func"
	"ojbk.io/gopherCron/cmd/service/controller/log_func"
	"ojbk.io/gopherCron/cmd/service/controller/project_func"
	"ojbk.io/gopherCron/cmd/service/controller/user_func"
	"ojbk.io/gopherCron/cmd/service/middleware"
)

func SetupRoute(r *gin.Engine) {
	r.Use(middleware.CrossDomain())
	r.Use(middleware.BuildResponse())

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.POST("/login", user_func.Login)
			user.Use(middleware.TokenVerify())
			user.GET("/info", user_func.GetUserInfo)
		}

		cron := api.Group("/crontab")
		{
			cron.Use(middleware.TokenVerify())
			cron.POST("/save", etcd_func.SaveTask)
			cron.POST("/delete", etcd_func.DeleteTask)
			cron.GET("/list", etcd_func.GetTaskList)
			cron.POST("/kill", etcd_func.KillTask)
			cron.GET("/worker_list", etcd_func.GetWorkerList)
		}

		project := api.Group("/project")
		{
			project.Use(middleware.TokenVerify())
			project.POST("/create", project_func.Create)
			project.GET("/list", project_func.GetUserProjects)
			project.POST("/delete", project_func.DeleteOne)
		}

		log := api.Group("/log")
		{
			log.Use(middleware.TokenVerify())
			log.GET("/list", log_func.GetList)
		}
	}
}