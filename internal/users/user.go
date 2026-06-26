package users

import (
	"strconv"
	"sync"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var mu sync.RWMutex

var all = []User{
	{
		ID:   "1",
		Name: "Gabriel",
	},
	{
		ID:   "2",
		Name: "valin",
	},
}

func List() []User {
	mu.RLock()
	defer mu.RUnlock()

	usersCopy := make([]User, len(all))
	copy(usersCopy, all)
	return usersCopy
}

func Create(req CreateUserRequest) User {
	mu.Lock()
	defer mu.Unlock()

	newUser := User{
		ID:    strconv.Itoa(len(all) + 1),
		Name:  req.Name,
		Email: req.Email,
	}

	all = append(all, newUser)

	return newUser
}
