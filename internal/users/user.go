package users

type User struct {
	ID   string
	Name string
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
