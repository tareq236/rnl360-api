package entity

import (
	DB "rnl360-api/database"
	"rnl360-api/models"
	u "rnl360-api/utils"
)

//Validate incoming user details...
func ValidateCelebration(celebration *models.CelebrationModel) (map[string]interface{}, bool) {

	if celebration.WorkArea == "" {
		return u.Message(false, "WorkArea is required", ""), false
	}

	if celebration.ChamberTerritoryID == "" {
		return u.Message(false, "ChamberTerritoryID is required", ""), false
	}

	return u.Message(false, "Requirement passed", ""), true
}

func SaveCelebration(celebration *models.CelebrationModel) map[string]interface{} {

	if resp, ok := ValidateCelebration(celebration); !ok {
		return resp
	}

	err_chk_duplicate := DB.GetDB().Where("work_area = ? and chamber_territory_id = ? and dr_child_id = ? and permission_status = ?", celebration.WorkArea, celebration.ChamberTerritoryID, celebration.DrChildID, 1).
		Last(&celebration).Error
	if err_chk_duplicate == nil {
		return u.Message(false, "Data already save.", "")
	}

	err := DB.GetDB().Create(celebration).Error
	if err != nil {
		return u.Message(false, "Failed to save.", err.Error())
	}

	err1 := DB.GetDB().Select("celebrations.*, text_message_list.details as text_message").
		Joins("LEFT JOIN text_message_list ON celebrations.text_message_id = text_message_list.id").
		Where("celebrations.id = ?", celebration.ID).First(&celebration).Error
	if err1 != nil {
		return u.Message(false, "Failed to fatch data.", err1.Error())
	}

	response := u.Message(true, "Celebration data has been saved", "")
	response["result"] = celebration
	return response
}

func CheckCelebrationPermission(celebration *models.CelebrationModel) (err error) {

	if err = DB.GetDB().Select("celebrations.*, text_message_list.details as text_message, response_type.response_type_name as permission_response_type_text").
		Joins("LEFT JOIN text_message_list ON celebrations.text_message_id = text_message_list.id").
		Joins("LEFT JOIN response_type ON celebrations.permission_response_type = response_type.id").
		Where("celebrations.work_area = ? and celebrations.chamber_territory_id = ? and celebrations.dr_child_id = ?", celebration.WorkArea, celebration.ChamberTerritoryID, celebration.DrChildID).
		Last(&celebration).Error; err != nil {
		return err
	}
	return nil

}

func UpdateCelebration(celebration *models.CelebrationModel) (err error) {
	if err = DB.GetDB().Model(&celebration).Where("id = ?", celebration.ID).Update(&celebration).Error; err != nil {
		return err
	}
	return nil
}

func CelebrationDetails(celebration *models.CelebrationModel) (err error) {
	if err = DB.GetDB().Select("celebrations.*, text_message_list.details as text_message").
		Joins("LEFT JOIN text_message_list ON celebrations.text_message_id = text_message_list.id").
		Where("celebrations.id = ?", celebration.ID).First(&celebration).Error; err != nil {
		return err
	}
	return nil
}

func GetAllCelebrationPendingRequest(celebration *[]models.CelebrationModel, work_area string) (err error) {

	if err = DB.GetDB().Select("celebrations.*, text_message_list.details as text_message, user_list.user_name, response_type.response_type_name as permission_response_type_text, (IF(celebrations.celebration_status = 1,'Celebrate','Do not celebrate')) as celebrate_status_text").
		Joins("LEFT JOIN text_message_list ON celebrations.text_message_id = text_message_list.id").
		Joins("LEFT JOIN user_list ON celebrations.work_area = user_list.work_area").
		Joins("LEFT JOIN response_type ON celebrations.permission_response_type = response_type.id").
		Where("celebrations.request_work_area = ? and celebrations.permission_status is NULL and celebrations.response_type = ?", work_area, 0).
		Order("celebrations.id DESC").
		Find(&celebration).Error; err != nil {
		return err
	}
	return nil

}

func GetAllCelebrationPermissionResponse(celebration *[]models.CelebrationModel, work_area string) (err error) {

	if err = DB.GetDB().Select("celebrations.*, text_message_list.details as text_message, user_list.user_name, response_type.response_type_name as permission_response_type_text").
		Joins("LEFT JOIN text_message_list ON celebrations.text_message_id = text_message_list.id").
		Joins("LEFT JOIN user_list ON celebrations.work_area = user_list.work_area").
		Joins("LEFT JOIN response_type ON celebrations.permission_response_type = response_type.id").
		Where("celebrations.request_work_area = ? and celebrations.permission_status IS NOT NULL and celebrations.response_type = ?", work_area, 0).
		Order("celebrations.id DESC").
		Find(&celebration).Error; err != nil {
		return err
	}
	return nil

}

func IsCelebrationChcekComplet(celebrationID int) (isComplet bool) {
	celebration := &models.CelebrationModel{}
	err := DB.GetDB().Where("id = ? and response_type = ?", celebrationID, 1).First(&celebration).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func IsCelebrationComplet(WorkArea string, DrChildID string) (isComplet bool) {
	celebration := &models.CelebrationModel{}
	err := DB.GetDB().Where("work_area = ? and dr_child_id = ? and response_type = ?", WorkArea, DrChildID, 1).Last(&celebration).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetAllCelebrationList(celebration *[]models.CelebrationModel, request_work_area string) (err error) {

	if err = DB.GetDB().Select("celebrations.*, text_message_list.details as text_message, user_list.user_name, response_type.response_type_name as permission_response_type_text").
		Joins("LEFT JOIN text_message_list ON celebrations.text_message_id = text_message_list.id").
		Joins("LEFT JOIN user_list ON celebrations.work_area = user_list.work_area").
		Joins("LEFT JOIN response_type ON celebrations.permission_response_type = response_type.id").
		Where("celebrations.request_work_area = ?", request_work_area).
		Order("celebrations.id DESC").
		Find(&celebration).Error; err != nil {
		return err
	}
	return nil

}
