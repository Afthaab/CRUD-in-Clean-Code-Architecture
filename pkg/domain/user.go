package domain

type User struct {
	ID   uint   `json:"id" gorm:"unique;not null"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}
