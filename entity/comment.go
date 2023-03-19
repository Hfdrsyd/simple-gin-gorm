package entity


type Comment struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	Text   string `json:"text"`
	UserID uint64 `json:"user_id"`
	User   *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	BlogID uint64 `gorm:"foreignKey" json:"blog_id"`
	Blog   *Blog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog"`
}