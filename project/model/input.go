package model

type RequestUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"`
}

type RequestOrders struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	NameProduct string `json:"name_product" binding:"required"`
	TotalOrder  int8   `json:"total_order" binding:"required"`
}
