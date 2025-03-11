package student

import "github.com/gin-gonic/gin"

func StudentRegister(router *gin.RouterGroup) {

	router.POST("/", AddStudent)
}
