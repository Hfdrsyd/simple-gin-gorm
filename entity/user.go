package entity

type User struct{
	ID uint64 `gorm:"primaryKey" json:"id"`
	Nama string `json:"nama"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
	Blogs []Blog `json:"blog,omitempty"`
	Likes []Like `json:"like,omitempty"`//melakukan like ke blog
	Comments []Comment  `json:"comment,omitempty"`
}
