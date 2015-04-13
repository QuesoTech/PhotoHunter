package main

type Feedback struct {
	Id            int64
	UserId        int64
	DatapointId   int64
	RequirementId int64
}

func NewFeedback(u_id, dp_id, req_id int64) (fb *Feedback, err error) {
	var id int64
	err = DB.QueryRow("INSERT INTO feedback (user_id, datapoint_id, requirement_id) VALUES ($1, $2, $3) RETURNING id", u_id, dp_id, req_id).Scan(&id)
	if err != nil {
		return
	}

	fb = &Feedback{
		Id:            id,
		UserId:        u_id,
		DatapointId:   dp_id,
		RequirementId: req_id,
	}

	return
}

func GetFeedbackById(id int64) (fb *Feedback, err error) {
	var u_id int64
	var dp_id int64
	var req_id int64
	err = DB.QueryRow("SELECT user_id, datapoint_id, requirement_id FROM feedback WHERE id=$1", id).Scan(&u_id, &dp_id, &req_id)
	if err != nil {
		return
	}

	fb = &Feedback{
		Id:            id,
		UserId:        u_id,
		DatapointId:   dp_id,
		RequirementId: req_id,
	}

	return
}
