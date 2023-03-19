package dto

type BlogCreate struct{
	ID uint64 `gorm:"primaryKey" json:"id"`
	Judul string `json:"judul" binding:"required"`
	Isi_blog string `json:"isi_blog" binding:"required"`
	User_ID uint64 `gorm:"foreignKey" json:"user_id"`
}
type BlogUpdate struct{
	ID uint64 `gorm:"primaryKey" json:"id"`
	Judul string `json:"judul" binding:"required"`
	Isi_blog string `json:"isi_blog" binding:"required"`
}