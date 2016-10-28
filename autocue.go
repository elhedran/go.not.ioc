package main

import (
	"time"
)

type TopgearAutoCue struct {
	Now func() time.Time
}

func (t *TopgearAutoCue) GetSomethingIntelligent() string {
	var now = t.now()
	if now.Weekday() == time.Tuesday {
		if now.Hour() >= 18 && now.Hour() < 20 {
			return "Good evening!"
		}
		return "Get out of the way!"
	}
	return "You're Wrong!"
}

func (t *TopgearAutoCue) now() time.Time {
	if t.Now == nil {
		return time.Now()
	}
	return t.Now()
}
