package models

import "time"

type BlogModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Details   string    `json:"details" gorm:"type:longtext"`
	Image     string    `json:"image"`
	Feature   int       `json:"feature"`
	Status    int       `gorm:"default:1"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *BlogModel) TableName() string {
	return "blogs"
}
