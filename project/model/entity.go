package model

type User struct {
	ID    int      `gorm:"PrimaryKey"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Order []Orders `gorm:"ForeignKey:IdUser"`
}

type Orders struct {
	ID         int    `gorm:"PrimaryKey"`
	IdUser     int8   `json:"id_user"`
	NameOrder  string `json:"name_order"`
	TotalOrder int8   `json:"total_order"`
}
