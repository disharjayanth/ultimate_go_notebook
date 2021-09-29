package users

type user struct {
	Name string
	ID   int
}

// user field isn't exported since it is lower case but user field with Name and ID
// can still be accessible from Manager since it is exported field in user type
// and those fields gets promoted in Manager type
type Manager struct {
	Title string
	user
}
