package models

import "time"

type UserModel struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	WorkArea          string    `json:"work_area"`
	UserName          string    `json:"user_name"`
	UserDesignationID int       `json:"user_designation_id"`
	Team              string    `json:"team"`
	PushKey           string    `json:"push_key"`
	DeviceType        int       `gorm:"default:1" json:"device_type"` // 1=Android; 2=iOS
	Status            int       `gorm:"default:1"`
	CreatedAt         time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *UserModel) TableName() string {
	return "user_list"
}
