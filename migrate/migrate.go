package main

import (
	"github.com/syahmiabbas/OneCV_Internship_test/initializers"
	"github.com/syahmiabbas/OneCV_Internship_test/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Teacher{}, &models.Student{})
	seed()
}

func seed() {
	type teachObj struct {
		Teachers models.Teacher
	}
	teachers := []models.Teacher{
		{Email: "teacherken@gmail.com"},
		{Email: "teacherken2@gmail.com"},
	}
	students := []models.Student{
		{Email: "studentjon@gmail.com", Suspended: false},
		{Email: "studenthon@gmail.com", Suspended: false},
		{Email: "studentagnes@gmail.com", Suspended: false},
		{Email: "studenmiche@gmail.com", Suspended: false},
	}

	initializers.DB.Create(&teachers)
	initializers.DB.Create(&students)
	initializers.DB.Save(&teachers)
	initializers.DB.Save(&students)
}
