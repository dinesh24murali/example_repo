package student

import (
	"github.com/gin-gonic/gin"
)

func StudentRegister(router *gin.RouterGroup) {

	studentService := NewStudentService()
	studentHandler := NewStudentHandler(studentService)

	router.POST("/", studentHandler.AddStudent)
	router.GET("/", studentHandler.GetStudents)
}
