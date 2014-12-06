package models

// Post model
type Post struct {
	ID      int64  `gorm:"column:ID; primary_key:yes"json:"id"`
	Content string `gorm:"column:post_content"json:"content"`
	Title   string `gorm:"column:post_title"json:"title"`
}

// TableName Method
func (p Post) TableName() string {
	return "gfb_posts"
}
