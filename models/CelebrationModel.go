package models

import (
	"time"
)

type CelebrationModel struct {
	ID                          uint      `gorm:"primaryKey" json:"id"`
	WorkArea                    string    `json:"work_area"`
	ChamberTerritoryID          string    `json:"chamber_territory_id"`
	DrChildID                   string    `json:"dr_child_id"`
	DoctorName                  string    `json:"doctor_name"`
	ChamberAddress              string    `gorm:"default:NULL" json:"chamber_address"`
	CellPhone                   string    `gorm:"default:NULL" json:"cell_phone"`
	Email                       string    `gorm:"default:NULL" json:"email"`
	DateOfBirth                 string    `json:"date_of_birth"`
	RequestWorkArea             string    `json:"request_work_area"`
	CelebrationType             int       `gorm:"default:NULL" json:"celebration_type"` // 1=Birthday; 2=Marragrday
	CelebrationStatus           int       `json:"celebration_status"`                   // Celebration=1; 0=Not Celebration;
	CelebrationCancelText       string    `gorm:"default:NULL" json:"celebration_cancel_text"`
	PermissionStatus            int       `gorm:"default:NULL" json:"permission_status"` // 1=ASK; 2=Permitted; 3=Cancle;
	PermissionRequestDateTime   time.Time `gorm:"default:NULL" json:"permission_request_date_time"`
	PermissionResponseDateTime  time.Time `gorm:"default:NULL" json:"permission_response_date_time"`
	PermissionResponseType      uint      `gorm:"default:NULL" json:"permission_response_type"`      // 0=Do Not Celebration; 1=Gift; 2=SMS; 3=EMAIL;
	PermissionResponseTypeText  string    `gorm:"default:NULL" json:"permission_response_type_text"` // GIFT,SMS,EMAIL
	PermissionResponseTypeEmail uint      `gorm:"default:0" json:"permission_response_type_email"`   // EMAIL; 0 not 1 yes
	PermissionResponseTypeSms   uint      `gorm:"default:0" json:"permission_response_type_sms"`     // SMS; 0 not 1 yes
	PermissionResponseTypeGift  uint      `gorm:"default:0" json:"permission_response_type_gift"`    // All type of gift; 0 not 1 yes
	PermissionResponseText      string    `gorm:"default:NULL" json:"permission_response_text"`
	ResponseDateTime            time.Time `gorm:"default:NULL" json:"response_date_time"`
	ResponseType                int       `gorm:"default:0" json:"response_type"` // 0=Not Complete; 1=Complete
	TextMessageID               int64     `gorm:"default:NULL" json:"text_message_id"`
	Picture                     string    `gorm:"default:NULL" json:"picture"`
	Feedback                    string    `gorm:"default:NULL" json:"feedback"`
	Status                      int       `gorm:"default:1"`
	CreatedAt                   time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt                   time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	TextMessage                 string    `sql:"-" json:"text_message"`
	UserName                    string    `sql:"-" json:"user_name"`
	ResponseTypeInput           string    `sql:"-" json:"response_type_input"`
	CelebrateStatusText         string    `sql:"-" json:"celebrate_status_text"`
	CelebrateStatus             string    `sql:"-" json:"celebrate_status"`
}

func (b *CelebrationModel) TableName() string {
	return "celebrations"
}
