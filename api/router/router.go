package router

import (
	"github.com/bqxtt/book_online/api/auth"
	"github.com/bqxtt/book_online/api/handler"
	"github.com/bqxtt/book_online/api/router/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

func Init() {
	router := gin.Default()
	router.Use(middleware.Cors(), middleware.Auth())
	user := router.Group("/user")
	{
		user.POST("/register", handler.Register)
		user.POST("/login", handler.Login)
		user.GET("/info", handler.GetUserInfo)
		user.POST("/info/update", handler.UpdateUserInfo)
	}

	admin := router.Group("/admin")
	{
		adminUser := admin.Group("/user")
		{
			adminUser.GET("/list", handler.ListAllUsers)
			adminUser.DELETE("/delete/:userId", handler.DeleteUser)
		}
		adminBook := admin.Group("/book")
		{
			adminBook.DELETE("/delete/:bookId", handler.DeleteBook)
			adminBook.POST("/update", handler.UpdateBook)
			adminBook.POST("/create", handler.CreateBook)
			bookRecord := adminBook.Group("/record")
			{
				bookRecord.GET("/borrow", handler.ListAllBorrowedBook)
				bookRecord.GET("/return", handler.ListAllReturnedBook)
				bookRecord.GET("/all", handler.ListAllBookRecords)
			}
		}
	}

	book := router.Group("/book")
	{
		book.GET("/list/:page/:pageSize", handler.ListBooks)
		bookRecord := book.Group("/record")
		{
			bookRecord.GET("/borrow", handler.ListBorrowedBook)
			bookRecord.GET("/return", handler.ListReturnedBook)
			bookRecord.GET("/all", handler.ListBookRecords)
		}
		book.GET("/borrow/:bookId", handler.BorrowBook)
		book.GET("/return/:bookId", handler.ReturnBook)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/token/admin", getAdminToken)
	router.GET("/token/user", getUserToken)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("router run error: %v", err)
	}
}

// generating jwt token for api test

// @Tags temp
// @Summary 生成管理员token
// @Description 生成管理员token
// @Accept  json
// @Produce json
// @Success 200 {string} string
// @Router /token/admin [get]
func getAdminToken(c *gin.Context) {
	token, err := auth.GenerateToken(10175101201, "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, token)
}

// @Tags temp
// @Summary 生成用户token
// @Description 生成用户token
// @Accept  json
// @Produce json
// @Success 200 {string} string
// @Router /token/user [get]
func getUserToken(c *gin.Context) {
	token, err := auth.GenerateToken(10175101201, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, token)
}
