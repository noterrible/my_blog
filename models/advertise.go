package models

type Advertise struct {
	Model
	Title  string `gorm:"title"`
	Href   string `gorm:"href"`
	IsShow bool   `gorm:"is_show"` //默认不展示
	Link   string `gorm:"link"`
}
