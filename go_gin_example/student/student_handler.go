package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service *StudentService
}

func NewStudentHandler(service *StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) AddStudent(c *gin.Context) {
	studentValidator := StudentValidator{}
	if err := studentValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := h.service.Create(&studentValidator.data)

	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	response := h.service.FindAll()
	c.JSON(http.StatusOK, gin.H{"data": response})
}
