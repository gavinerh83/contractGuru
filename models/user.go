package user

//DepartmentMap maps out the role to an integer for easier database access
var DepartmentMap = map[string]int{
	"business requester": businessRequester,
	"owner":              owner,
	"finance":            finance,
	"legal":              legal,
	"contract admin":     contractAdministrator,
}

//Data constructs the user information
type Data struct {
	UserID     int
	Firstname  string
	Lastname   string
	Email      string
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
