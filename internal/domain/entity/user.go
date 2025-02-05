package entity

type User struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Education string `json:"education"`
}

type CreateUserDTO struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Education string `json:"education"`
}

type UpdateUserDTO struct {
	Name      string `json:"name" `
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Education string `json:"education"`
}
