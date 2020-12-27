package handler

import (
	"github.com/bqxtt/book_online/api/model/contract"
	"github.com/bqxtt/book_online/api/service"
	"github.com/bqxtt/book_online/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 注册
// @Description 用户注册
// @Accept  json
// @Produce json
// @Param   request body contract.RegisterRequest true "register request"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /register [post]
func Register(c *gin.Context) {
	request := &contract.RegisterRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailureResponse("request error, err: %v", err))
		return
	}
	if err := service.UserService.Register(request.User); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailureResponse("user service error, err: %v", err))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("register success"))
}

func Login(c *gin.Context) {

}
