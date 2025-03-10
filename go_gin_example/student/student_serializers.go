package student

import (
	"github.com/gin-gonic/gin"
)

type StudentSerializer struct {
	c *gin.Context
}

type Student struct {
	ID          string `json:"id"`
	FirstName   string `json:"title"`
	LastName    string `json:"description"`
	RollNo      string `json:"roll_no"`
	Mark        int    `json:"mark"`
	DateOfBirth string `json:"dob"`
}
