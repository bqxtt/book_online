package handler

import (
	"fmt"
	"github.com/bqxtt/book_online/api/auth"
	"github.com/bqxtt/book_online/api/model/contract"
	"github.com/bqxtt/book_online/api/service"
	"github.com/bqxtt/book_online/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags user
// @Summary 借阅图书
// @Description 借阅图书
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.BorrowBookRequest true "borrow book request"
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
	if request.BookId <= 0 {
		c.JSON(http.StatusBadRequest, &contract.BorrowBookResponse{
			BaseResponse: utils.NewFailureResponse("book id is invalid, bookId=[%v]", request.BookId),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	userId := claims.UserId
	fmt.Println(userId)
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.BorrowBookResponse{
			BaseResponse: utils.NewFailureResponse("user id error, err: %v", err),
		})
		return
	}
	deadline, err := service.RentService.BorrowBook(id, request.BookId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.BorrowBookResponse{
			BaseResponse: utils.NewFailureResponse("rent service error, err: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, &contract.BorrowBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Deadline:     deadline.Format("2006/01/02"),
	})
}

// @Tags user
// @Summary 归还图书
// @Description 归还图书
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ReturnBookRequest true "return book request"
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
	err := service.RentService.ReturnBook(request.RecordId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ReturnBookResponse{
			BaseResponse: utils.NewFailureResponse("rent service error, err: %v", err),
		})
		return
	}
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
// @Param request body contract.ListBorrowedBookRequest true "list borrowed book request"
// @Success 200 {object} contract.ListBorrowedBookResponse
// @Failure 400 {object} contract.ListBorrowedBookResponse
// @Router /book/record/borrow [post]
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
	id, _ := strconv.ParseInt(userId, 10, 64)
	request := &contract.ListBorrowedBookRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListBorrowedBookResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	records, pageInfo, err := service.RentService.ListBorrowedBook(id, request.Page, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListBorrowedBookResponse{
			BaseResponse: utils.NewFailureResponse("rent service error, err: %v", err),
			Records:      nil,
			PageInfo:     nil,
		})
		return
	}

	c.JSON(http.StatusOK, &contract.ListBorrowedBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      records,
		PageInfo:     pageInfo,
	})
}

// @Tags user
// @Summary 已还图书记录
// @Description 已还图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param request body contract.ListReturnedBookRequest true "list returned book request"
// @Success 200 {object} contract.ListReturnedBookResponse
// @Failure 400 {object} contract.ListReturnedBookResponse
// @Router /book/record/return [post]
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
	id, _ := strconv.ParseInt(userId, 10, 64)
	request := &contract.ListReturnedBookRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListReturnedBookResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	records, pageInfo, err := service.RentService.ListReturnedBook(id, request.Page, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListReturnedBookResponse{
			BaseResponse: utils.NewFailureResponse("rent service error, err: %v", err),
			Records:      nil,
			PageInfo:     nil,
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ListReturnedBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      records,
		PageInfo:     pageInfo,
	})
}

// @Tags user
// @Summary 所有借阅图书记录
// @Description 所有借阅图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param request body contract.ListBookRecordsRequest true "list book records request"
// @Success 200 {object} contract.ListBookRecordsResponse
// @Failure 400 {object} contract.ListBookRecordsResponse
// @Router /book/record/all [post]
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
	id, _ := strconv.ParseInt(userId, 10, 64)
	request := &contract.ListBookRecordsRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListBookRecordsResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	records, pageInfo, err := service.RentService.ListBook(id, request.Page, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListBookRecordsResponse{
			BaseResponse: utils.NewFailureResponse("rent service error, err: %v", err),
			Records:      nil,
			PageInfo:     nil,
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ListBookRecordsResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      records,
		PageInfo:     pageInfo,
	})
}

// @Tags admin
// @Summary 所有用户待还图书记录
// @Description 所有用户待还图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param request body contract.ListAllBorrowedBookRequest true "list all borrowed book records request"
// @Success 200 {object} contract.ListAllBorrowedBookResponse
// @Failure 400 {object} contract.ListAllBorrowedBookResponse
// @Router /admin/book/record/borrow [post]
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

	request := &contract.ListAllBorrowedBookRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListAllBorrowedBookResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}

	records, pageInfo, err := service.RentService.ListAllBorrowedBook(request.Page, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListAllBorrowedBookResponse{
			BaseResponse: utils.NewFailureResponse("rent service error, err:%v", err),
			Records:      nil,
			PageInfo:     nil,
		})
		return
	}

	c.JSON(http.StatusOK, &contract.ListAllBorrowedBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      records,
		PageInfo:     pageInfo,
	})
}

// @Tags admin
// @Summary 所有用户已还图书记录
// @Description 所有用户已还图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param request body contract.ListAllReturnedBookRequest true "list all returned book records request"
// @Success 200 {object} contract.ListAllReturnedBookResponse
// @Failure 400 {object} contract.ListAllReturnedBookResponse
// @Router /admin/book/record/return [post]
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
	request := &contract.ListAllReturnedBookRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListAllReturnedBookResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	records, pageInfo, err := service.RentService.ListAllReturnedBook(request.Page, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListAllReturnedBookResponse{
			BaseResponse: utils.NewFailureResponse("rent service error, err: %v", err),
			Records:      nil,
			PageInfo:     nil,
		})
		return
	}

	c.JSON(http.StatusOK, &contract.ListAllReturnedBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      records,
		PageInfo:     pageInfo,
	})
}

// @Tags admin
// @Summary 所有用户所有借阅图书记录
// @Description 所有用户所有借阅图书记录
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param request body contract.ListAllBookRecordsRequest true "list all book records request"
// @Success 200 {object} contract.ListAllBookRecordsResponse
// @Failure 400 {object} contract.ListAllBookRecordsResponse
// @Router /admin/book/record/all [post]
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
	request := &contract.ListAllBookRecordsRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListAllBookRecordsResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	records, pageInfo, err := service.RentService.ListAllBookRecords(request.Page, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListAllBookRecordsResponse{
			BaseResponse: utils.NewFailureResponse("rent service error, err: %v", err),
			Records:      nil,
			PageInfo:     nil,
		})
		return
	}
	c.JSON(http.StatusOK, &contract.ListAllBookRecordsResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Records:      records,
		PageInfo:     pageInfo,
	})
}
