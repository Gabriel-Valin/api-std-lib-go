package users

import "sync"

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Mu sync.RWMutex

var All = []User{
	{
		ID:   "1",
		Name: "Gabriel",
	},
	{
		ID:   "2",
		Name: "valin",
	},
}
