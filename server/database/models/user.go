package models

type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"size:100;not null" json:"name"`
	Email string `gorm:"size:100;unique" json:"email"`
	
}
