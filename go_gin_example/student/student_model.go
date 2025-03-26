package student

type Student struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	RollNo      string `json:"roll_no"`
	DateOfBirth string `json:"dob"`
	Mark        int    `json:"mark"`
}
