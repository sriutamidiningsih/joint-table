package model

type ResponseUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseJoin struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	NameProduct string `json:"name_product"`
	TotalOrder  int8   `json:"total_order"`
}