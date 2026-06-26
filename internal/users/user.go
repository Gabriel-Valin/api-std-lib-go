package users

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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
