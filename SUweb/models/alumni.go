package models

type Alumni struct {
	AlumniId uint `gorm:"primary_key" json:"alumni_id" gorm:"auto-increment"`
	Name string `json:"name"`
	Mail string `json:"mail"`
	City string `json:"city"`
	Phone string `json:"phone"`
	Type string `json:"type"`
	GraduateYear string `json:"graduate_year"`
	Company string `json:"company"`
	SchoolP string `json:"schoolP"`
	Wishes string `json:"wishes"`
	Introduction string `json:"introduction"`
}

func QueryAlumni(alumnis *[]Alumni) error {
	if err := db.Find(&alumnis).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (alumni *Alumni) CreateAlumni() error {
	if db.Where("name = ?",alumni.Name).Error!=nil {
		return db.Update(alumni).Error
	}else {
		return db.Create(&alumni).Error
	}
}

func (alumni *Alumni) QueryAlumni(username string) error {
	return db.Where("name = ?",username).First(&alumni).Error
}
