package handler

import (
	"github.com/bqxtt/book_online/api/model/contract"
	"github.com/bqxtt/book_online/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	request := &contract.RegisterRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := service.UserService.Register(request.User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "register success",
	})
}

func Login(c *gin.Context) {

}
