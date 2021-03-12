package models

//DepartmentMap maps out the role to an integer for easier database access
var DepartmentMap = map[string]int{
	"business requester": businessRequester,
	"owner":              owner,
	"finance":            finance,
	"legal":              legal,
	"contract admin":     contractAdministrator,
}

var (
	//SessionMap stores the session information for the user
	SessionMap = map[string]string{} //JWT as the key, value as the username
	//UserMap stores signed up user information
	UserMap = map[string]Data{} //store user information, key is username, value is userinfo created in signup
)

//Data constructs the user information
type Data struct {
	UserID     int
	Firstname  string
	Lastname   string
	Username   string //or the email
	Password   string
	Department int
}

const (
	businessRequester = iota
	owner
	finance
	legal
	contractAdministrator
)
