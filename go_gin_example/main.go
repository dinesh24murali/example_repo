package main

import (
	"go_gin_example/student"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	student.StudentRegister(r.Group("/student"))

	r.Run(":8080")
}
