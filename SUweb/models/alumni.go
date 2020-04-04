package models

type Alumni struct {
	AlumniId uint `gorm:"primary_key" json:"alumni_id" gorm:"auto-increment"`
	Name string `json:"name"`
	Introduction string `json:"introduction"`
	AuthId uint `gorm:"ForeignKey:AuthId"`
}

func QueryAlumni(alumnis *[]Alumni) error {
	if err := db.Find(&alumnis).Error; err != nil {
		return err
	} else {
		return nil
	}
}
