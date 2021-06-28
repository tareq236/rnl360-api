package entity

import (
	DB "rnl360-api/database"
	"rnl360-api/models"
)

func GetAllResponseType(responseType *[]models.ResponseTypeModel) (err error) {

	if err = DB.GetDB().Where("status = ?", 1).Order("id ASC").Find(&responseType).Error; err != nil {
		return err
	}
	return nil

}

func GetResponseTypeName(responseType *models.ResponseTypeModel, responseTypeName string) (err error) {

	if err = DB.GetDB().Where("response_type_name = ?", responseTypeName).First(&responseType).Error; err != nil {
		return err
	}
	return nil

}
