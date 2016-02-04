package main

type IDAL interface {
	CreateUser(APIUser) (APIUser, error)
	FindUser(APIUser) (APIUser, error)
	UpdateUser(APIUser) (APIUser, error)
	DeleteUser(APIUser) (APIUser, error)
}
