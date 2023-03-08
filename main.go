package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syahmiabbas/OneCV_Internship_test/controllers"
	"github.com/syahmiabbas/OneCV_Internship_test/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/api/register", controllers.TeachersCreate)
	r.GET("/api/commonstudents", controllers.CommonStudents)
	r.POST("/api/suspend", controllers.SuspendStudent)
	r.POST("/api/retrievefornotifications", controllers.RetrieveFornotification)

	r.Run()
}
