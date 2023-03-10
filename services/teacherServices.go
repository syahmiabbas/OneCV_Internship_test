package services

import (
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/syahmiabbas/OneCV_Internship_test/initializers"
	"github.com/syahmiabbas/OneCV_Internship_test/models"
)

func TeachersCreateService(teacher string, students []string) struct {
	Code    int
	Message string
} {
	type structResponse struct {
		Code    int
		Message string
	}
	Teacher, rowsAffected := checkExistingTeacher(teacher)
	if rowsAffected == 0 {

		return structResponse{
			Code:    412,
			Message: "Teacher `" + teacher + "` does not exist",
		}
	}

	var studentEmails []models.Student
	for _, student := range students {
		_, rowsAffectedStudent := checkExistingStudent(student)

		if rowsAffectedStudent != 0 {
			studentObj := models.Student{Email: student}
			studentEmails = append(studentEmails, studentObj)
		} else {
			return structResponse{
				Code:    412,
				Message: "Student `" + student + "` does not exist",
			}
		}
	}

	initializers.DB.Model(&Teacher).Association("Students").Append(&studentEmails)
	return structResponse{
		Code:    204,
		Message: "",
	}
}

func CommonStudentsService(teacherList []string) struct {
	Code    int
	Message string
	Data    interface{}
} {
	var Students []models.Student

	type structResponse struct {
		Code    int
		Message string
		Data    interface{}
	}

	for _, teacher := range teacherList {
		_, rowsAffected := checkExistingTeacher(teacher)
		if rowsAffected == 0 {
			return structResponse{
				Code:    412,
				Message: "Teacher `" + teacher + "` does not exist",
				Data:    nil,
			}
		}
	}

	initializers.DB.Where("email IN (?)", initializers.DB.Table("teacher_student").
		Select("student_email").
		Where("teacher_email IN (?) AND suspended = ?", teacherList, false).
		Group("student_email").Having("COUNT(student_email) = ?", len(teacherList)),
	).Find(&Students)

	return structResponse{
		Code:    200,
		Message: "",
		Data:    Students,
	}

}

func SuspendStudentService(student string) struct {
	Code    int
	Message string
} {
	type structResponse struct {
		Code    int
		Message string
	}

	Student, rowsAffectedStudent := checkExistingStudent(student)
	if rowsAffectedStudent == 0 {

		return structResponse{
			Code:    412,
			Message: "Student `" + student + "` does not exist",
		}
	}

	Student.Suspended = true
	initializers.DB.Save(&Student)

	return structResponse{
		Code:    204,
		Message: "",
	}
}

func RetrieveFornotificationService(teacher string, notification string) struct {
	Code    int
	Message string
	Data    interface{}
} {
	type structResponse struct {
		Code    int
		Message string
		Data    interface{}
	}
	_, rowsAffected := checkExistingTeacher(teacher)
	if rowsAffected == 0 {
		return structResponse{
			Code:    412,
			Message: "Teacher `" + teacher + "` does not exist",
			Data:    nil,
		}
	}

	notificationList := strings.Fields(notification)
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
		Where("teacher_email IN (?) AND suspended = ?", teacher, false),
	).Find(&RegisteredStudents)

	for _, student := range RegisteredStudents {
		recipientList.Add(student.Email)
	}

	return structResponse{
		Code:    200,
		Message: "",
		Data:    recipientList,
	}
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
