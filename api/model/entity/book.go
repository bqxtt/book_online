package entity

type Book struct {
	BookId      int64  `form:"book_id" json:"book_id"`
	ImgUrl      string `form:"img_url" json:"img_url"`
	BookName    string `form:"book_name" json:"book_name" binding:"required"`
	ISBN        string `form:"isbn" json:"isbn" binding:"required"`
	Publisher   string `form:"publisher" json:"publisher" binding:"required"`
	Category    string `form:"category" json:"category"`
	Author      string `form:"author" json:"author" binding:"required"`
	Description string `form:"description" json:"description"`
	//1：未借出 2：已借出
	RentStatus int8 `form:"rent_status" json:"rent_status"`
}
