package entity

type User struct {
	ID        string        `json:"id"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Address   []UserAddress `json:"address"`
}

type UserAddress struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}
