package model

type userModel struct{}

var UserModel userModel

type User struct{
	Id int64 `json:"id"`
	
}