package entity

type User struct {
	Id           uint64 `json:"id" gorm:"primary_key; auto_increment"`
	Name         string `json:"name" gorm:"type:varchar(100)"`
	Password     string `json:"password" gorm:"type:varchar(100)"`
	Age          int    `json:"age" binding:"gte=10,lte=70" gorm:"type:int(2)"`
	Email        string `json:"email" binding:"required,email" gorm:"type:varchar(150)"`
	ImageUrl     string `json:"imageUrl" binding:"url" gorm:"type:varchar(255)"`
	Status       string `json:"status" gorm:"type:varchar(50)"`
	SocketId     string `json:"socketId" gorm:"type:varchar(100)"`
	RefreshToken string `json:"refreshToken" gorm:"type:varchar(255)"`
}
