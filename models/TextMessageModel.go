package models

import "time"

type TextMessageModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Details   string    `json:"details" gorm:"type:longtext"`
	Type      int       `gorm:"default:1" json:"type"` // 1=sms; 2=email;
	Status    int       `gorm:"default:1"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *TextMessageModel) TableName() string {
	return "text_message_list"
}
