package entity
type Blog struct{
	ID uint64   `gorm:"primaryKey" json:"id"`
	Judul string `json:"judul"`
	Isi_blog string `json:"isi_blog"`
	User_ID uint64 `gorm:"foreignKey" json:"user_id"`
	User *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"buat_blog"`
	Likes []Like `json:"likes,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}