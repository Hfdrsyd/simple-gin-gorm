package entity

type Authorization struct {
	Token string `gorm:"type:varchar(255)" json:"token"`
	Role  string `gorm:"type:varchar(30)" json:"role"`
}