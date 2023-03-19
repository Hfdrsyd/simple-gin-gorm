package dto

type RegisterUser struct{
	ID uint64 `gorm:"primaryKey" json:"id"`
	Nama string `json:"nama" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role string `json:"role" binding:"required"`
}
type LoginUser struct{
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UpdateName struct {
	Nama string `json:"nama" binding:"required"`
}
