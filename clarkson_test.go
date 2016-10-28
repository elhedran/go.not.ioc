package main

import (
	"testing"
)

type alwaysHello struct{}

func (a *alwaysHello) GetSomethingIntelligent() string {
	return "hello"
}

type rememberLast struct {
	last string
}

func (r *rememberLast) Shout(something string) {
	r.last = something
}

func TestClarksonReads(t *testing.T) {
	remember := &rememberLast{}
	always := &alwaysHello{}
	clarkson := &Clarkson{
		AutoCue:   always,
		Megaphone: remember,
	}
	clarkson.ShoutSomething()

	actual := remember.last
	expected := always.GetSomethingIntelligent()
	if actual != expected {
		t.Errorf("actual: %q, expected: %q\n", actual, expected)
		t.Fail()
	}
}
