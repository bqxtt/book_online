package handler

import "github.com/gin-gonic/gin"

// @Tags user
// @Summary 借阅图书
// @Description 借阅图书
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   book_id path int true "borrow book id"
// @Success 200 {object} contract.BorrowBookResponse
// @Failure 400 {object} contract.BorrowBookResponse
// @Router /book/borrow/{book_id} [get]
func BorrowBook(c *gin.Context) {

}

// @Tags user
// @Summary 归还图书
// @Description 归还图书
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication Token"
// @Param   book_id path int true "return book id"
// @Success 200 {object} contract.ReturnBookResponse
// @Failure 400 {object} contract.ReturnBookResponse
// @Router /book/return/{book_id} [get]
func ReturnBook(c *gin.Context) {

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

}
