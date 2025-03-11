package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var students = []Student{}
var nextID = 2

func AddStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
