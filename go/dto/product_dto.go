package dto

type ProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Desc  string `json:"desc"`
	Price int64  `json:"price" binding:"required"`
}
