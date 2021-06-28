package entity

import (
	DB "rnl360-api/database"
	"rnl360-api/models"
)

func SaveUser(user *models.UserModel) (err error) {
	if err = DB.GetDB().Create(user).Error; err != nil {
		return err
	}
	return nil
}

func CheckUser(user *[]models.UserModel, workArea string) (err error) {
	if err = DB.GetDB().Where("work_area = ?", workArea).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.UserModel) (err error) {
	// DB.GetDB().Where("work_area = ?", user.WorkArea).First(&user)
	// db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	if err = DB.GetDB().Model(&user).Where("work_area = ?", user.WorkArea).Update(&user).Error; err != nil {
		return err
	}
	return nil
}
