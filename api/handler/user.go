package handler

import (
	"github.com/bqxtt/book_online/api/model/contract"
	"github.com/bqxtt/book_online/api/service"
	"github.com/bqxtt/book_online/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags user
// @Summary 注册
// @Description 用户注册
// @Accept  json
// @Produce json
// @Param   request body contract.RegisterRequest true "register request"
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.RegisterResponse
// @Failure 400 {object} contract.RegisterResponse
// @Router /user/register [post]
func Register(c *gin.Context) {
	request := &contract.RegisterRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.RegisterResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	if err := service.UserService.Register(request.UserAuth); err != nil {
		c.JSON(http.StatusBadRequest, &contract.RegisterResponse{
			BaseResponse: utils.NewFailureResponse("user service error, err: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, &contract.RegisterResponse{
		BaseResponse: utils.NewSuccessResponse("register success"),
	})
}

// @Tags user
// @Summary 登录
// @Description 用户登录
// @Accept  json
// @Produce json
// @Param   request body contract.LoginRequest true "login request"
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.LoginResponse
// @Failure 400 {object} contract.LoginResponse
// @Router /user/login [post]
func Login(c *gin.Context) {
	request := &contract.LoginRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.LoginResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	if err := service.UserService.Login(request.UserAuth); err != nil {
		c.JSON(http.StatusBadRequest, &contract.LoginResponse{
			BaseResponse: utils.NewFailureResponse("user service error, err: %v", err),
		})
		return
	}
}

// @Tags user
// @Summary 获取用户信息
// @Description 用户信息
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.GetUserInfoResponse
// @Failure 400 {object} contract.GetUserInfoResponse
// @Router /user/info [get]
func GetUserInfo(c *gin.Context) {

}

// @Tags user
// @Summary 更新用户信息
// @Description 更新用户信息
// @Accept  json
// @Produce json
// @Param   request body contract.UpdateUserInfoRequest true "update user info request"
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.UpdateUserInfoResponse
// @Failure 400 {object} contract.UpdateUserInfoResponse
// @Router /user/info/update [post]
func UpdateUserInfo(c *gin.Context) {

}

// @Tags admin
// @Summary 获取所有用户
// @Description 所有用户
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.ListAllUsersResponse
// @Failure 400 {object} contract.ListAllUsersResponse
// @Router /admin/user/list [get]
func ListAllUsers(c *gin.Context) {

}

// @Tags admin
// @Summary 删除用户
// @Description 删除用户
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   user_id path int true "delete user id"
// @Success 200 {object} contract.DeleteUserResponse
// @Failure 400 {object} contract.DeleteUserResponse
// @Router /admin/user/delete/{user_id} [delete]
func DeleteUser(c *gin.Context) {

}
