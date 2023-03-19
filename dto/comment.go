package dto

type Comment struct{
	ID uint64 `gorm:"primaryKey" json:"id"`
	Text      string `json:"text" binding:"required"`
}