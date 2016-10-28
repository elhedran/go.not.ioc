package main

type AutoCue interface {
	GetSomethingIntelligent() string
}

type Megaphone interface {
	Shout(string)
}
