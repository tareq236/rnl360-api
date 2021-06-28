package models

type LoginModel struct {
	UserCode          string `json:"user_code"`
	UserName          string `json:"user_name"`
	UserDesignationID int    `json:"user_designation_id"`
	WorkAreaT         string `json:"work_area_t"`
	WorkArea          string `json:"work_area"`
	Team              string `json:"team"`
	Password          string `json:"password"`
	PushKey           string `json:"push_key" sql:"-"`
	DeviceType        int    `json:"device_type" sql:"-"`
}

func (b *LoginModel) TableName() string {
	return "UserLogin"
}
