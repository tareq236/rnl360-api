package entity

import (
	DB "rnl360-api/database"
	"rnl360-api/models"
)

func GetAllCommunication(communicationModel *[]models.CommunicationModel) (err error) {

	if err = DB.GetDB().Where("status = ?", 1).Order("id DESC").Find(&communicationModel).Error; err != nil {
		return err
	}
	return nil

}
