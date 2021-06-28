package models

import "time"

type ResponseTypeModel struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	ResponseTypeName string    `json:"response_type_name"`
	Status           int       `gorm:"default:1"`
	CreatedAt        time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *ResponseTypeModel) TableName() string {
	return "response_type"
}
