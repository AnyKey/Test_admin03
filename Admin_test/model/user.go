package model

type User struct {
	Iduser   int    `json:"iduser"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Userlog struct {
	User_id   int    `json:"user_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Token string `json:"token"`
	Expire_time string `json:"expire_time"`
}
type TokenTest struct {
	User_id int `json:"user_id"`
	Token string `json:"token"`
	Expire_time string `json:"expire_time"`
}
type ErrorTest struct {
	ErrorT string `json:"error_t"`
}
type Test int

type Authorization bool