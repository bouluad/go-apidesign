package models

type User struct {
	FirstName string `json:"username"`
	LastName  string `json:"lastname"`
}

func NewUser() *User {
	return &User{}
}

type UserModel struct{}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (this UserModel) GetUsers() User {
	u := User{FirstName: "Mohammed", LastName: "BOULUAD"}

	return u
}
