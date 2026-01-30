package models

import "gorm.io/gorm"

type Anime struct {
	gorm.Model        //otomatis membuat id,created_at dan updated_at -> delete_at
	Title      string `json:"title"`
	Studio     string `json:"studio"`
	Rating     int    `json:"rating"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqueIndex"` // ini buat supaya email itu uniq
	Password string `json:"-"`                        //supaya tidak tampil di response json
}
