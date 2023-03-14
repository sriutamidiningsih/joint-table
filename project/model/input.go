package model

type RequestUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"`
}

type RequestJoin struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	NameOrder  string `json:"name_order"`
	TotalOrder int8   `json:"total_order" binding:"required"`
}
