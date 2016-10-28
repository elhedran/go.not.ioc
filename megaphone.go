package main

import (
	"github.com/fatih/color"
)

type ConsoleMegaphone struct{}

func (m *ConsoleMegaphone) Shout(something string) {
	color.Red(something)
}
