package main

type Clarkson struct {
	AutoCue   AutoCue
	Megaphone Megaphone
}

func (c *Clarkson) ShoutSomething() {
	c.megaphone().Shout(c.autoCue().GetSomethingIntelligent())
}

var DefaultMegaphone = &ConsoleMegaphone{}
var DefaultAutoCue = &TopgearAutoCue{}

func (c *Clarkson) megaphone() Megaphone {
	if c.Megaphone == nil {
		return DefaultMegaphone
	}
	return c.Megaphone
}

func (c *Clarkson) autoCue() AutoCue {
	if c.AutoCue == nil {
		return DefaultAutoCue
	}
	return c.AutoCue
}
