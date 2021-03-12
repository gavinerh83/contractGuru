package storage

//contains functions for interacting with the database for contract related requests

const (
	//add in the table names as constants
	tableUser                 = "user"
	tableUserRole             = "user_role"
	tableRequest              = "request"
	tableRequestTransition    = "request_transition"
	tableState                = "state"
	tableTransition           = "transition"
	tableAction               = "action"
	tableTransitionAction     = "transition_action"
	tableEvent                = "event"
	tableTransitionEvent      = "transition_event"
	tableCompletedRequestData = "completed_request_data"
	tableCompletedRequestFile = "completed_request_file"
)

type user struct {
	userID    int
	username  string
	password  string
	firstname string
	lastname  string
}

type userRole struct {
	userRoleID   int
	userRoleName string
}

type request struct {
	requestID      int
	requesterID    int
	buID           int
	requestName    string
	requestDate    string
	financeFlag    int
	currentStateID int
}

type requestTransition struct {
	requestID    int
	transitionID int
	userID       int
	doneDate     string
}

type state struct {
	stateID   int
	stateName string
}

type transition struct {
	transitionID int
	userRoleID   int
	startStateID int
	endStateID   int
}

type action struct {
	actionID   int
	actionName string
}

type transitionAction struct {
	transitionID int
	actionID     int
}

type event struct {
	eventID   int
	eventName string
}

type transitionEvent struct {
	transitionID int
	eventID      int
}

type completedRequestData struct {
	completedID             int
	requestID               int
	counterPartyName        string
	counterPartyInformation string
	contractValue           int
	contractType            string //??
	region                  string
	effectiveDate           string
	terminationDate         string
	renewalDate             string
	purpose                 string
	background              string
}

type completedRequestFile struct {
	fileID      int
	completedID int
	uploadDate  string
	fileContent string
	userID      int
	mimeType    string
}
