package entity

import (
	DB "rnl360-api/database"
	"rnl360-api/models"
)

func GetAllBlogFeature(blogModel *[]models.BlogModel) (err error) {

	if err = DB.GetDB().Where("feature = ?", 1).Order("id ASC").Find(&blogModel).Error; err != nil {
		return err
	}
	return nil

}
