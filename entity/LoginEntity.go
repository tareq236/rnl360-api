package entity

import (
	DB "rnl360-api/database"
)

type LoginResult struct {
	RequestUserCode   string `json:"request_user_code"`
	UserCode          string `json:"user_code"`
	UserName          string `json:"user_name"`
	UserDesignationID int    `json:"user_designation_id"`
	WorkAreaT         string `json:"work_area_t"`
	WorkArea          string `json:"work_area"`
	Team              string `json:"team"`
}

func Login(workArea string, password string, userResult *[]LoginResult) (err error) {

	if err = DB.GetSQLDB().Raw("SELECT ULV.UserCode AS RequestUserCode, UL.* FROM RDB.dbo.UserLogin UL LEFT JOIN UserLevel ULV ON UL.WorkArea = ULV.AssignUserCode WHERE UL.WorkArea = ? AND Password = ?;", workArea, password).Scan(userResult).Error; err != nil {
		return err
	}
	return nil

}
