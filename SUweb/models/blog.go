package models

import "time"

type Blog struct {
	BlogId         uint                 `gorm:"primary_key" json:"blog_id" gorm:"auto-increment"`
	CreatedAt      time.Time            `json:"created_at"`
	Title          string               `json:"title"`
	AuthorId       uint                 `json:"author_id"`
	AuthorName     string               `json:"author_name"`
	ContentHtml    string               `json:"content_html"`
	ContentMd      string               `json:"content_md"`
	ReadTimes      int                  `json:"read_times" gorm:"default:0;"`
	StarTimes      int                  `json:"star_times" gorm:"default:0;"`
	ThumbsTimes    int                  `json:"thumbs_times" gorm:"default:0;"`
	ChangeTimes    int                  `json:"change_times" gorm:"default:0;"`
	Comments       []BlogComment        `json:"comments"`
	DataStatistics []BlogDataStatistics `json:"data_statistics"`
}

type BlogComment struct {
	CommentId string    `json:"comment_id"`
	AuthorId  int       `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type BlogDataStatistics struct {
	BlogStatisticsId  uint `json:"blog_statistics_id" gorm:"primary_key;"`
	BlogId            uint `json:"blog_id" gorm:"ForeignKey:BlogId"`
	ReaderId          uint `json:"reader_id" gorm `
	ReadTime          uint `json:"read_time"`
	FinalStayPosition int  `json:"final_stay_position"`
}

type WatchLater struct {
	id     uint `gorm:"primary_key" gorm:"auto-increment"`
	UserId uint `json:"user_id"`
	BlogId uint `json:"blog_id"`
}

func (wl *WatchLater) Create() error {
	return db.Create(&wl).Error
}

func QueryBlogList(blogs *[]Blog) error {
	//var projects []Project
	if err := db.Find(&blogs).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func QueryBlogListByUserId(blogs *[]Blog, userid uint) error {
	//var projects []Project
	if err := db.Where("author_id = ?", userid).Find(&blogs).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (blog *Blog) QueryBlog(id string) error {
	return db.Where("blog_id = ?", id).First(&blog).Error
}

func (blog *Blog) CreateBlog() error {
	return db.Create(&blog).Error
}

func (blog *Blog) UpdateBlog() error {
	return db.Update(&blog).Error
}
func QueryBlogComments(comments *[]BlogComment, project_id string) error {
	return db.Where("blog_id = ?", project_id).Find(&comments).Error
}

func QueryWatchLater(wls *[]WatchLater, user_id string) error {
	return db.Where("user_id = ?", user_id).Find(&wls).Error
}

func QueryBlogNums() int64 {
	var result int64
	db.Count(&result)
	return result
}
