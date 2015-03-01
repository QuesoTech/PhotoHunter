package main

import (
	"time"
)

type TimePeriod struct {
	Id        int
	DatasetId int
	Start     time.Time
	End       time.Time
}
