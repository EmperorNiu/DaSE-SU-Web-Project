package models

import "time"

type Project struct {
	ProjectId          uint     `gorm:"primary_key" json:"project_id" gorm:"auto-increment"`
	ProjectName        string   `json:"project_name" form:"project_name" binding:"required"`
	ProjectLeader      string   `json:"project_leader" form:"project_leader" binding:"required"`
	ProjectDescription string   `form:"project_description"`
	Members            []Member `gorm:"many2many:project_members"`
	Tutor              Tutor    `gorm:"ForeignKey:TutorId"`
	TutorId            int
	Comments           []ProjectComment
}

type ProjectIntro struct {
	ProjectId     uint     `json:"proj_id "`
	ProjectLeader string   `json:"leader" `
	Members       []Member `json:"members"`
	ProjectName   string   `json:"proj_name"`
}

type Member struct {
	MemberId       uint
	MemberName     string
	Responsibility string
	Auth           Auth `gorm:"ForeignKey:AuthId"`
	AuthId         int
}

type Tutor struct {
	TutorId   uint
	TutorName string
}

type ProjectComment struct {
	CommentId uint   `gorm:"primary_key;auto-increment" json:"comment_id"`
	Auth      string `gorm:"ForeignKey:AuthId"`
	AuthId    int
	Content   string    `json:"content"`
	ProjectId int       `json:"project_id" gorm:"ForergnKey:ProjectId"`
	CreatedAt time.Time `json:"created_at"`
}

func QueryProjectList(projects *[]Project) error {
	//var projects []Project
	if err := db.Find(&projects).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (project *Project) Create() error {
	err := db.Create(&project).Error
	return err
}

func (project *Project) Insert() error {
	err := db.Create(&project).Error
	return err
}

func (project *Project) Update() error {
	return db.Save(&project).Error
}

func (project *Project) QueryProject(id string) error {
	return db.Where("project_id = ?", id).First(&project).Error
}

func (comment *ProjectComment) CreateComment() error {
	return db.Create(&comment).Error
}

func QueryComments(comments *[]ProjectComment, project_id string) error {
	return db.Where("project_id = ?", project_id).Find(&comments).Error
}
