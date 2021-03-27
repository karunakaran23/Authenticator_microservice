package model

type User struct {
	ID       uint32 `json:"id"`
	Username string `json:"name"`
}
type UserRequest struct {
	Username string `json:"name"`
}
