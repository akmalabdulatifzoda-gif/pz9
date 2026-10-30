package core

import "time"

// User — пользователь. Пароль хранится только в виде bcrypt-хэша
// и никогда не попадает в JSON (тег json:"-").
type User struct {
	ID           int64     `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;size:255;not null" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
