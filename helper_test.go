package fuzzyfinder

func New() *Finder {
	return &Finder{}
}

func NewWithMockedTerminal() (*Finder, *TerminalMock) {
	f := New()
	m := f.UseMockedTerminal()
	w, h := 60, 10 // A normally value.
	m.SetSize(w, h)
	return f, m
}
