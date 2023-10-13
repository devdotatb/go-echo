package service

type ServiceInterface interface {
	GetHello() string
	UpdateHello(text string)
	GetAllUser() []User
	AddUser(u User)
}
