package entity

type Board struct {
	ID            string `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	OrgID         string `gorm:"size:100;not null" json:"org_id"`
	Title         string `gorm:"size:255;not null" json:"title"`
	ImageID       string `gorm:"size:100" json:"image_id"`
	ImageThumbUrl string `gorm:"size:255" json:"image_thumb_url"`
	ImageFullUrl  string `gorm:"size:255" json:"image_full_url"`
	ImageUserName string `gorm:"size:100" json:"image_user_name"`
	ImageLinkHtml string `gorm:"size:255" json:"image_link_html"`
}
