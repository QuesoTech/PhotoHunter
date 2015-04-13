package main

import (
	"time"
)

// TimePeriod represents the time restriction placed on a dataset (i.e.
// can only be taken at night)
type TimePeriod struct {
	Id        int
	DatasetId int
	Start     time.Time
	End       time.Time
}
