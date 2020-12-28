package handler

import "github.com/gin-gonic/gin"

// @Tags user
// @Summary 图书列表
// @Description 分页图书
// @Accept  json
// @Produce json
// @Param   page path int true "page"
// @Param   page_size path int true "page size"
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.ListBooksResponse
// @Failure 400 {object} contract.ListBooksResponse
// @Router /book/list/{page}/{page_size} [get]
func ListBooks(c *gin.Context) {

}

// @Tags admin
// @Summary 新增图书
// @Description 新增图书
// @Accept  json
// @Produce json
// @Param   request body contract.CreateBookRequest true "create book request"
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.CreateBookResponse
// @Failure 400 {object} contract.CreateBookResponse
// @Router /admin/book/create [post]
func CreateBook(c *gin.Context) {

}

// @Tags admin
// @Summary 更改图书信息
// @Description 更改图书
// @Accept  json
// @Produce json
// @Param   request body contract.UpdateBookRequest true "update book request"
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.UpdateBookResponse
// @Failure 400 {object} contract.UpdateBookResponse
// @Router /admin/book/update [post]
func UpdateBook(c *gin.Context) {

}

// @Tags admin
// @Summary 删除图书
// @Description 删除图书
// @Accept  json
// @Produce json
// @Param   book_id path int true "delete book id"
// @Param Authorization header string true "Authentication Token"
// @Success 200 {object} contract.DeleteBookResponse
// @Failure 400 {object} contract.DeleteBookResponse
// @Router /admin/book/delete/{book_id} [delete]
func DeleteBook(c *gin.Context) {

}
