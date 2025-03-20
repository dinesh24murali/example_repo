package student

import (
	"github.com/gin-gonic/gin"
)

type StudentValidator struct {
	FirstName   string `form:"first_name" json:"first_name" binding:"required,min=2,max=155"`
	LastName    string `form:"last_name" json:"last_name" binding:"required,min=2,max=255"`
	RollNo      string `form:"roll_no" json:"roll_no" binding:"required,min=3,max=3"`
	Mark        int    `form:"mark" json:"mark" binding:"required,gte=0,lte=100"`
	DateOfBirth string `form:"dob" json:"dob" binding:"required,date"`
	// DateOfBirth string  `form:"dob" json:"dob"`
	data Student `json:"-"`
}

func (p *StudentValidator) Bind(c *gin.Context) error {
	if err := c.BindJSON(p); err != nil {
		return err
	}
	p.data.FirstName = p.FirstName
	p.data.LastName = p.LastName
	p.data.RollNo = p.RollNo
	p.data.Mark = p.Mark
	p.data.DateOfBirth = p.DateOfBirth
	return nil
}
