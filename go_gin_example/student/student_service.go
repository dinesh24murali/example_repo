package student

var students = []Student{}
var nextID = 1

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (r *StudentService) Create(student *Student) *Student {
	student.ID = nextID
	nextID++

	students = append(students, *student)

	return student
}

func (r *StudentService) FindAll() []Student {
	return students
}
