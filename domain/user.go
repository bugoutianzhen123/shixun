package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `json:"username"`
	Password   string `json:"password"`
	Permission uint   `json:"permission" gorm:"default:1"` // 0 :high    1 :low
}
