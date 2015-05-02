package main

import (
	"time"
)

// TimePeriod represents the time restriction placed on a dataset (i.e.
// can only be taken at night)
type TimePeriod struct {
	ID        int64
	DatasetID int64
	Start     time.Time
	End       time.Time
}

func NewTimeperiod(dsID int64, start, end time.Time) (t *TimePeriod, err error) {
	var id int64

	err = DB.QueryRow("INSERT INTO subject (ds_id, start, end) VALUES ($1) RETURNING id", dsID, start, end).Scan(&id)
	if err != nil {
		return
	}

	t = &TimePeriod{
		ID:        id,
		DatasetID: dsID,
		Start:     start,
		End:       end,
	}

	return
}
