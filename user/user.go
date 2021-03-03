package user

var userMap map[string]User

type User struct {
	firstname  string
	lastname   string
	email      string
	password   string
	department int
}
