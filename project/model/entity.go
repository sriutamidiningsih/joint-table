package model

type User struct {
	ID      int       `gorm:"PrimaryKey"`
	Name    string    `json:"name"`
	NoHp    int       `json:"nohp"`
	UserAdress []Address `gorm:"foreignKey:UserID"`
}

type Address struct {
	ID     int `gorm:"PrimaryKey"`
	UserID int
	Kota   string `json:"kota"`
	Negara string `json:"negara"`
}
