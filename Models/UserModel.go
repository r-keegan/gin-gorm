package Models

type User struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	PinNumber int    `json:"pin"`
}

func (b *User) TableName() string {
	return "user"
}
