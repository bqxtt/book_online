package handler

import (
	"fmt"
	"github.com/bqxtt/book_online/api/auth"
	"github.com/bqxtt/book_online/api/model/contract"
	"github.com/bqxtt/book_online/api/service"
	"github.com/bqxtt/book_online/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags user
// @Summary 图书列表
// @Description 分页图书
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.ListBooksRequest true "list book request"
// @Success 200 {object} contract.ListBooksResponse
// @Failure 400 {object} contract.ListBooksResponse
// @Router /book/list [post]
func ListBooks(c *gin.Context) {
	_, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.ListBooksResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	request := &contract.ListBooksRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListBooksResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}
	if request.Page < 1 || request.PageSize < 0 {
		c.JSON(http.StatusBadRequest, &contract.ListBooksResponse{
			BaseResponse: utils.NewFailureResponse("page or page size is incorrect"),
		})
	}
	books, totalPages, err := service.BookService.ListBooksByPage(request.Page, request.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, &contract.ListBooksResponse{
			BaseResponse: utils.NewFailureResponse("book service is error, err: %v", err),
		})
	}
	c.JSON(http.StatusOK, &contract.ListBooksResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
		Books:        books,
		TotalPages:   totalPages,
	})
}

// @Tags admin
// @Summary 新增图书
// @Description 新增图书
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.CreateBookRequest true "create book request"
// @Success 200 {object} contract.CreateBookResponse
// @Failure 400 {object} contract.CreateBookResponse
// @Router /admin/book/create [post]
func CreateBook(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.CreateBookResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	if claims.Role != "admin" {
		c.JSON(http.StatusBadRequest, &contract.CreateBookResponse{
			BaseResponse: utils.NewFailureResponse("not admin"),
		})
		return
	}
	request := &contract.CreateBookRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.CreateBookResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
	}

	c.JSON(http.StatusOK, &contract.CreateBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
	})
}

// @Tags admin
// @Summary 更改图书信息
// @Description 更改图书
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   request body contract.UpdateBookRequest true "update book request"
// @Success 200 {object} contract.UpdateBookResponse
// @Failure 400 {object} contract.UpdateBookResponse
// @Router /admin/book/update [post]
func UpdateBook(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.UpdateBookResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	if claims.Role != "admin" {
		c.JSON(http.StatusBadRequest, &contract.UpdateBookResponse{
			BaseResponse: utils.NewFailureResponse("not admin"),
		})
		return
	}
	request := &contract.UpdateBookRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.UpdateBookResponse{
			BaseResponse: utils.NewFailureResponse("request param error, err: %v", err),
		})
	}

	c.JSON(http.StatusOK, &contract.UpdateBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
	})
}

// @Tags admin
// @Summary 删除图书
// @Description 删除图书
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   book_id path int true "delete book id"
// @Success 200 {object} contract.DeleteBookResponse
// @Failure 400 {object} contract.DeleteBookResponse
// @Router /admin/book/delete/{book_id} [delete]
func DeleteBook(c *gin.Context) {
	iClaims, exist := c.Get("auth")
	if !exist {
		c.JSON(http.StatusBadRequest, &contract.DeleteBookResponse{
			BaseResponse: utils.NewFailureResponse("missing auth header"),
		})
		return
	}
	claims := iClaims.(*auth.Claims)
	if claims.Role != "admin" {
		c.JSON(http.StatusBadRequest, &contract.DeleteBookResponse{
			BaseResponse: utils.NewFailureResponse("not admin"),
		})
		return
	}
	bookId := c.Param("bookId")
	fmt.Println(bookId)

	c.JSON(http.StatusOK, &contract.DeleteBookResponse{
		BaseResponse: utils.NewSuccessResponse("success"),
	})
}
