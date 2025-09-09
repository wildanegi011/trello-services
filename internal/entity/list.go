package entity

type List struct {
	ID        string `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	Title     string `gorm:"size:255;not null" json:"title"`
	Position  int    `gorm:"not null" json:"position"`
	BoardID   string `gorm:"size:100;not null;index;column:boardId" json:"board_id"`
	Board     Board  `gorm:"foreignKey:BoardID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
	CreatedAt string `gorm:"autoCreateTime;column:createdAt" json:"created_at"`
	UpdatedAt string `gorm:"autoUpdateTime;column:updatedAt" json:"updated_at"`
}
