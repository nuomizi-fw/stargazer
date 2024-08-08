package model

type User struct {
	BaseModel
	Username string `json:"username"`
}
