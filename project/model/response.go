package model

type ResponseUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseJoin struct {
	ID         int    `json:"id"`
	NameOrder  string `json:"name_order"`
	TotalOrder int8   `json:"total_order"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}
