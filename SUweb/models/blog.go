package models

import "time"

type Blog struct {
	BlogId int `gorm: "primary_key" json: "id"`
	CreatedAt time.Time `json:"created_at"`
	Title string `json:"title"`
	AuthorId int `json:"author_id"`
	AuthorName string `json:"author_name"`
	ContentHtml string `json:"content_html"`
	ContentMd string `json:"content_md"`
	ReadTimes int `json:"read_times"`
	StarTimes int `json:"star_times"`
	ThumbsTimes int `json:"thumbs_times"`
	ChangeTimes int `json:"change_times"`
	Comments []BlogComment `json:"comments"`
	DataStatistics []BlogDataStatistics `json:"data_statistics"`
}

type BlogComment struct {
	CommentId string `json:"comment_id"`
	AuthorId int `json:"author_id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type BlogDataStatistics struct {
	BlogStatisticsId uint `json:"blog_statistics_id" gorm:"primary_key;"`
	BlogId uint `json:"blog_id" gorm:"ForeignKey:BlogId"`
	ReaderId uint `json:"reader_id" gorm `
	ReadTime uint `json:"read_time"`
	FinalStayPosition int `json:"final_stay_position"`
}

func QueryBlogList(blogs *[]Blog) error {
	//var projects []Project
	if err := db.Find(&blogs).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (blog *Blog) QueryBlog(id string) error {

}
