package main

import (
	"go_gin_example/student"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {

	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("date", student.ValidateDate)
	}

	student.StudentRegister(r.Group("/student"))

	r.Run(":8080")
}
