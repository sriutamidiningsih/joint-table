package model

type Users struct {
	ID        int      `gorm:"PrimaryKey"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	OrderList []Orders `gorm:"ForeignKey:IdUser"`
}

type Orders struct {
	ID          int    `gorm:"PrimaryKey"`
	IdUser      int8   `json:"id_user"`
	NameProduct string `json:"name_product"`
	TotalOrder  int8   `json:"total_order"`
	
}
