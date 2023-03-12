package model

type RequestUser struct {
	Name   string `json:"name" binding:"required"`
	NoHp   int    `json:"nohp" binding:"required,number"`
}
