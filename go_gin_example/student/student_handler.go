package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var students = []Student{}
var nextID = 2

func CreateTodo(c *gin.Context) {
	var todo Student
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
