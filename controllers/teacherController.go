package controllers

import (
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gin-gonic/gin"
	"github.com/syahmiabbas/OneCV_Internship_test/initializers"
	"github.com/syahmiabbas/OneCV_Internship_test/models"
)

func TeachersCreate(c *gin.Context) {
	var body struct {
		Teacher  string
		Students []string
	}
	c.Bind(&body)

	Teacher, rowsAffected := checkExistingTeacher(body.Teacher)
	if rowsAffected == 0 {
		c.JSON(412, gin.H{
			"message": "Teacher `" + body.Teacher + "` does not exist",
		})
		return
	}

	var studentEmails []models.Student
	for _, student := range body.Students {
		_, rowsAffectedStudent := checkExistingStudent(student)

		if rowsAffectedStudent != 0 {
			studentObj := models.Student{Email: student}
			studentEmails = append(studentEmails, studentObj)
		} else {
			c.JSON(412, gin.H{
				"message": "Student `" + student + "` does not exist",
			})
		}
	}

	initializers.DB.Model(&Teacher).Association("Students").Append(&studentEmails)

	c.JSON(204, gin.H{
		"message": "Successfully registered students",
	})

}

func CommonStudents(c *gin.Context) {

	teachers := c.Request.URL.Query()
	var Students []models.Student

	if teacherList, found := teachers["teacher"]; found {
		for _, teacher := range teacherList {

			_, rowsAffected := checkExistingTeacher(teacher)
			if rowsAffected == 0 {
				c.JSON(412, gin.H{
					"message": "Teacher `" + teacher + "` does not exist",
				})
				return
			}
			// initializers.DB.Model(&Teacher).Association("Students").Find(&Student)

		}
		// initializers.DB.Preload("Students").Where("email IN (SELECT student_email FROM teacher_student WHERE teacher_email IN ?)", teacherList).
		// 	Find(&Students)

		initializers.DB.Where("email IN (?)", initializers.DB.Table("teacher_student").
			Select("student_email").
			Where("teacher_email IN (?) AND suspended = ?", teacherList, false),
		).Find(&Students)

		c.JSON(200, gin.H{
			"students": Students,
		})
		return

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

	Student, rowsAffectedStudent := checkExistingStudent(body.Student)
	if rowsAffectedStudent == 0 {
		c.JSON(412, gin.H{
			"message": "Student `" + body.Student + "` does not exist",
		})
	}

	Student.Suspended = true
	initializers.DB.Save(&Student)

	c.JSON(204, gin.H{})
}

func RetrieveFornotification(c *gin.Context) {
	var body struct {
		Teacher      string
		Notification string
	}
	c.Bind(&body)

	_, rowsAffected := checkExistingTeacher(body.Teacher)
	if rowsAffected == 0 {
		c.JSON(412, gin.H{
			"message": "Teacher `" + body.Teacher + "` does not exist",
		})
		return
	}

	notificationList := strings.Fields(body.Notification)
	recipientList := mapset.NewSet[string]()
	var RegisteredStudents []models.Student

	for _, notification := range notificationList {
		if notification[0:1] == "@" {
			studentStr := notification[1:]
			_, rowsAffectedStudent := checkExistingNonSuspendedStudent(studentStr)

			if rowsAffectedStudent != 0 {
				recipientList.Add(notification[1:])
			}
		}
	}

	initializers.DB.Where("email IN (?)", initializers.DB.Table("teacher_student").
		Select("student_email").
		Where("teacher_email IN (?) AND suspended = ?", body.Teacher, false),
	).Find(&RegisteredStudents)

	for _, student := range RegisteredStudents {
		recipientList.Add(student.Email)
	}

	c.JSON(200, gin.H{
		"recipients": recipientList,
	})
}

func checkExistingStudent(student string) (models.Student, int64) {
	var Student models.Student
	checkStudent := initializers.DB.First(&Student, "email = ?", student)
	return Student, checkStudent.RowsAffected
}

func checkExistingNonSuspendedStudent(student string) (models.Student, int64) {
	var Student models.Student
	checkStudent := initializers.DB.First(&Student, "email = ? AND suspended = ?", student, false)
	return Student, checkStudent.RowsAffected
}

func checkExistingTeacher(teacher string) (models.Teacher, int64) {
	var Teacher models.Teacher
	checkTeacher := initializers.DB.First(&Teacher, "email = ?", teacher)
	return Teacher, checkTeacher.RowsAffected
}
