package users

import (
	"sync"
)

type MemoryStore struct {
	mu   sync.RWMutex
	data []User
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: []User{
			{
				ID:    1,
				Name:  "Gabriel",
				Email: "gabriel@email.com",
			},
			{
				ID:    2,
				Name:  "Maria",
				Email: "maria@email.com",
			},
		},
	}
}

func (s *MemoryStore) List() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	usersCopy := make([]User, len(s.data))
	copy(usersCopy, s.data)

	return usersCopy
}

// func (s *MemoryStore) Create(req CreateUserRequest) User {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	user := User{
// 		ID:    strconv.Itoa(len(s.data) + 1),
// 		Name:  req.Name,
// 		Email: req.Email,
// 	}

// 	s.data = append(s.data, user)

// 	return user
// }

func (s *MemoryStore) GetByID(id string) (User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.data {
		if user.ID == 1 {
			return user, true
		}
	}

	return User{}, false
}
