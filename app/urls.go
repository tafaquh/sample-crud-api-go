package app

import (
	"os"

	"github.com/gin-gonic/gin"

	"coba-go/controller"
)

func mapUrls() {

	v1 := router.Group("api/v1")
	{
		userPath := v1.Group("/user", gin.BasicAuth(gin.Accounts{
			os.Getenv("SERVICE_USERNAME"): os.Getenv("SERVICE_PASSWORD"),
		}))
		{
			registerPath := userPath.Group("/register")
			{
				registerPath.POST("/", controller.RegisterUser)
				//registerPath.POST("/teacher", controller.RegisterAdmin)
				//registerPath.POST("/student", controller.RegisterAdmin)
			}
			loginPath := userPath.Group("/login")
			{
				loginPath.POST("/", controller.Login)
			}
		}
		coursePath := v1.Group("/course", gin.BasicAuth(gin.Accounts{
			os.Getenv("SERVICE_USERNAME"): os.Getenv("SERVICE_PASSWORD"),
		}))
		{
			coursePath.POST("/", controller.CreateCourse)
			coursePath.GET("/:id", controller.GetCourseByID)
			coursePath.POST("/upload", controller.CourseFileUpload)
		}
	}

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
