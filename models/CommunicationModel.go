package models

import "time"

type CommunicationModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Details   string    `json:"details" gorm:"type:longtext"`
	FIleType  string    `json:"file_type"`
	FIle      string    `json:"file"`
	Type      string    `json:"type"`
	Zone      string    `json:"zone"`
	Status    int       `gorm:"default:1"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *CommunicationModel) TableName() string {
	return "communication"
}
