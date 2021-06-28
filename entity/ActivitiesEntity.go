package entity

import (
	DB "rnl360-api/database"
	"strconv"
	"time"
)

type ActivitiesResult struct {
	ID                    string `json:"id"`
	DrMasterID            string `json:"dr_master_id"`
	DrChildID             string `json:"dr_child_id"`
	DoctorName            string `json:"doctor_name"`
	Ch_Addr               string `json:"ch_addr"`
	CellPhone1            string `json:"cell_phone"`
	Email1                string `json:"email"`
	ProfessionalDegrees   string `json:"professional_degrees"`
	SpecialityDescription string `json:"speciality_description"`
	DOB                   string `json:"dob"`
	DOM                   string `json:"dom"`
}

func GetTodayActivities(workArea string, activitiesResult *[]ActivitiesResult) (err error) {

	_, month, day := time.Now().Date()
	currentMonthDay := strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	if err = DB.GetSQLDB().Raw("SELECT CT.ID, DM.DrMasterID, DC.DrChildID, RTRIM(DM.DoctorName1 + ' ' + DM.DoctorName2) AS DoctorName, DM.CellPhone1, DM.Email1, DM.ProfessionalDegrees, S.SpecialityDescription, RTRIM(DC.Ch_Addr1 + ' ' + DC.Ch_Addr2) AS Ch_Addr, DM.DOB, DM.DOM FROM dbo.ChamberTerritory CT  INNER JOIN DoctorsChamberP DC ON CT.DrChildID=DC.DrChildID INNER JOIN DoctorsMasterP DM ON DC.DrMasterID=DM.DrMasterID INNER JOIN Speciality S ON DC.SpecialityCode=S.SpecialityCode WHERE CT.WorkAreaT = ? AND (DM.DOB LIKE '%"+currentMonthDay+"' OR DM.DOM LIKE '%"+currentMonthDay+"');", workArea).Scan(activitiesResult).Error; err != nil {
		return err
	}
	return nil

}

func GetUpcomingActivities(workArea string, activitiesResult *[]ActivitiesResult) (err error) {

	if err = DB.GetSQLDB().Raw("SELECT CT.ID, DM.DrMasterID, DC.DrChildID, RTRIM(DM.DoctorName1 + ' ' + DM.DoctorName2) AS DoctorName, DM.CellPhone1, DM.Email1, DM.ProfessionalDegrees, S.SpecialityDescription, RTRIM(DC.Ch_Addr1 + ' ' + DC.Ch_Addr2) AS Ch_Addr, DM.DOB, DM.DOM FROM dbo.ChamberTerritory CT  INNER JOIN DoctorsChamberP DC ON CT.DrChildID=DC.DrChildID INNER JOIN DoctorsMasterP DM ON DC.DrMasterID=DM.DrMasterID INNER JOIN Speciality S ON DC.SpecialityCode=S.SpecialityCode WHERE CT.WorkAreaT = ? AND LEFT(CONVERT(VARCHAR(15), DM.DOB, 110), 5) BETWEEN LEFT(CONVERT(VARCHAR(15), GETDATE(), 110), 5) AND LEFT(CONVERT(VARCHAR(15), GETDATE()+70, 110), 5);", workArea).Scan(activitiesResult).Error; err != nil {
		return err
	}
	return nil

}
