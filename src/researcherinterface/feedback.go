package main

// Feedback is an instance of a user indicating whether a datapoint fulfills
// a requirement. Feedback should be unique per user, datapoint, requirement
// triple (i.e. no user should provide feedback on the same datapoint-requrement)
// pair
type Feedback struct {
	ID int64

	// id of providing user
	UserID int64

	// id of datapoint feedback is being provided for
	DatapointID int64

	// id of requirement being fed-back(?) on
	RequirementID int64
}

// NewFeedback creates a new feedback struct and inserts it to the database
func NewFeedback(uID, dpID, reqID int64) (fb *Feedback, err error) {
	var id int64
	err = DB.QueryRow("INSERT INTO feedback (userID, datapointID, requirementID) VALUES ($1, $2, $3) RETURNING id", uID, dpID, reqID).Scan(&id)
	if err != nil {
		return
	}

	fb = &Feedback{
		ID:            id,
		UserID:        uID,
		DatapointID:   dpID,
		RequirementID: reqID,
	}

	return
}

// GetFeedbackByID fetechs a feedback from the database
func GetFeedbackByID(id int64) (fb *Feedback, err error) {
	var uID int64
	var dpID int64
	var reqID int64
	err = DB.QueryRow("SELECT userID, datapointID, requirementID FROM feedback WHERE id=$1", id).Scan(&uID, &dpID, &reqID)
	if err != nil {
		return
	}

	fb = &Feedback{
		ID:            id,
		UserID:        uID,
		DatapointID:   dpID,
		RequirementID: reqID,
	}

	return
}
