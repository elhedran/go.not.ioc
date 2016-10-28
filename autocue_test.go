package main

import (
	"testing"
	"time"
)

func TestTuesdayEvening(t *testing.T) {
	cue := &TopgearAutoCue{}
	cue.Now = func() time.Time {
		return time.Date(
			2016, time.October, 25,
			19, 0, 0, 0,
			time.Local,
		)
	}
	actual := cue.GetSomethingIntelligent()
	expected := "Good evening!"
	if actual != expected {
		t.Errorf("actual: %q, expected: %q\n", actual, expected)
		t.Fail()
	}
}
