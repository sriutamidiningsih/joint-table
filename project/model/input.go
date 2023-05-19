package model

type RequestUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"`
}

type RequestOrders struct {
	IdUser      int8   `json:"id_user"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	NameProduct string `json:"name_product" binding:"required"`
	TotalOrder  int8   `json:"total_order" binding:"required"`
}
