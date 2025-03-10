package student

type StudentValidator struct {
	FirstName   string `form:"title" json:"title" binding:"required,min=2,max=155"`
	LastName    string `form:"description" json:"description" binding:"required,min=2,max=255"`
	RollNo      string `form:"roll_no" json:"roll_no" binding:"required,min=3,min=3"`
	Mark        int    `form:"mark" json:"mark" binding:"required,min=0,min=100"`
	DateOfBirth string `form:"dob" json:"dob" binding:"required,datetime"`
}

func NewTodoValidator() StudentValidator {
	todoValidator := StudentValidator{}
	return todoValidator
}
