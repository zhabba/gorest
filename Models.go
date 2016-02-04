package main

type APIUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type APIMessage struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
