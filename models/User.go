package models

type User struct{
	UserID 	int `json:"user_id"`
	Name	string	`json:"name"`
	Email	string	`json:"email"`
	Password	string	`json:"password"`
	Gender		string	`json:"gender"`
	Birthday	string	`json:"birthday"`
	CreatedAt	string	`json:"created_at"`
	Avatar		string	`json:"avatar_url"`
}