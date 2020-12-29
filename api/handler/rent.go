package handler

import (
	"fmt"
	"github.com/bqxtt/book_online/api/auth"
	"github.com/bqxtt/book_online/api/model/contract"
	"github.com/bqxtt/book_online/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags user
// @Summary 借阅图书
// @Description 借阅图书
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request {object} contract.BorrowBookRequest true "borrow book request"
// @Success 200 {object} contract.BorrowBookResponse
// @Failure 400 {object} contract.BorrowBookResponse
// @Router /book/borrow [post]
func BorrowBook(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.BorrowBookResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	request := &contract.BorrowBookRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.BorrowBookResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	userId := claims.UserId
	fmt.Println(userId)

	c.JSON(http.StatusOK, &contract.BorrowBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
	})
}

// @Tags user
// @Summary 归还图书
// @Description 归还图书
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request {object} contract.ReturnBookRequest true "return book request"
// @Success 200 {object} contract.ReturnBookResponse
// @Failure 400 {object} contract.ReturnBookResponse
// @Router /book/return [post]
func ReturnBook(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.ReturnBookResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	request := &contract.ReturnBookRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ReturnBookResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	userId := claims.UserId
	fmt.Println(userId)

	c.JSON(http.StatusOK, &contract.ReturnBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
	})
}

// @Tags user
// @Summary 待还图书记录
// @Description 待还图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.ListBorrowedBookResponse
// @Failure 400 {object} contract.ListBorrowedBookResponse
// @Router /book/record/borrow [get]
func ListBorrowedBook(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.ListBorrowedBookResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	userId := claims.UserId
	fmt.Println(userId)

	c.JSON(http.StatusOK, &contract.ListBorrowedBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      nil,
	})
}

// @Tags user
// @Summary 已还图书记录
// @Description 已还图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.ListReturnedBookResponse
// @Failure 400 {object} contract.ListReturnedBookResponse
// @Router /book/record/return [get]
func ListReturnedBook(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.ListReturnedBookResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	userId := claims.UserId
	fmt.Println(userId)

	c.JSON(http.StatusOK, &contract.ListReturnedBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      nil,
	})
}

// @Tags user
// @Summary 所有借阅图书记录
// @Description 所有借阅图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.ListBookRecordsResponse
// @Failure 400 {object} contract.ListBookRecordsResponse
// @Router /book/record/all [get]
func ListBookRecords(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.ListBookRecordsResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	userId := claims.UserId
	fmt.Println(userId)

	c.JSON(http.StatusOK, &contract.ListBookRecordsResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      nil,
	})
}

// @Tags admin
// @Summary 所有用户待还图书记录
// @Description 所有用户待还图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.ListAllBorrowedBookResponse
// @Failure 400 {object} contract.ListAllBorrowedBookResponse
// @Router /admin/book/record/borrow [get]
func ListAllBorrowedBook(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.ListAllBorrowedBookResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	if claims.Role != "admin" {
		c.JSON(http.StatusBadRequest, &contract.ListAllBorrowedBookResponse{
			BaseResponse: utils.NewFailureResponse("not admin"),
		})
		return
	}

	c.JSON(http.StatusOK, &contract.ListAllBorrowedBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      nil,
	})
}

// @Tags admin
// @Summary 所有用户已还图书记录
// @Description 所有用户已还图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.ListAllReturnedBookResponse
// @Failure 400 {object} contract.ListAllReturnedBookResponse
// @Router /admin/book/record/return [get]
func ListAllReturnedBook(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.ListAllReturnedBookResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	if claims.Role != "admin" {
		c.JSON(http.StatusBadRequest, &contract.ListAllReturnedBookResponse{
			BaseResponse: utils.NewFailureResponse("not admin"),
		})
		return
	}

	c.JSON(http.StatusOK, &contract.ListAllReturnedBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      nil,
	})
}

// @Tags admin
// @Summary 所有用户所有借阅图书记录
// @Description 所有用户所有借阅图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.ListAllBookRecordsResponse
// @Failure 400 {object} contract.ListAllBookRecordsResponse
// @Router /admin/book/record/all [get]
func ListAllBookRecords(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.ListAllBookRecordsResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	if claims.Role != "admin" {
		c.JSON(http.StatusBadRequest, &contract.ListAllBookRecordsResponse{
			BaseResponse: utils.NewFailureResponse("not admin"),
		})
		return
	}

	c.JSON(http.StatusOK, &contract.ListAllBookRecordsResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      nil,
	})
}
