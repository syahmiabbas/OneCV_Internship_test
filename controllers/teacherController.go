package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/syahmiabbas/OneCV_Internship_test/services"
)

func TeachersCreate(c *gin.Context) {
	var body struct {
		Teacher  string
		Students []string
	}
	c.Bind(&body)

	response := services.TeachersCreateService(body.Teacher, body.Students)

	c.JSON(response.Code, gin.H{
		"message": response.Message,
	})

}

func CommonStudents(c *gin.Context) {

	teachers := c.Request.URL.Query()

	if teacherList, found := teachers["teacher"]; found {

		response := services.CommonStudentsService(teacherList)
		if response.Code == 200 {
			c.JSON(response.Code, gin.H{
				"students": response.Data,
			})
		} else {
			c.JSON(response.Code, gin.H{
				"message": response.Message,
			})
		}

	} else {
		c.JSON(412, gin.H{
			"message": "No 'teacher' URL parameter was provided",
		})
	}
}

func SuspendStudent(c *gin.Context) {

	var body struct {
		Student string
	}
	c.Bind(&body)

	if body.Student == "" {
		c.JSON(412, gin.H{
			"message": "No student was provided",
		})
		return
	}

	response := services.SuspendStudentService(body.Student)

	c.JSON(response.Code, gin.H{
		"message": response.Message,
	})
}

func RetrieveFornotification(c *gin.Context) {
	var body struct {
		Teacher      string
		Notification string
	}
	c.Bind(&body)

	response := services.RetrieveFornotificationService(body.Teacher, body.Notification)
	if response.Code == 200 {
		c.JSON(response.Code, gin.H{
			"recipients": response.Data,
		})
	} else {
		c.JSON(response.Code, gin.H{
			"message": response.Message,
		})
	}

}
