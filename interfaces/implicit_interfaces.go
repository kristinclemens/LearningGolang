package interfaces

// An example of implicit interfaces. TabbyCat implements Cat and HouseCat implicitly because it has functions matching
// the signature of both interfaces. Lion implements only Cat because it does NOT have a function matching the HouseCat
// signature.
// Ref: https://dev.to/mraszplewicz/golang-through-the-eyes-of-a-java-developer-pros-and-cons-25o2

type Cat interface {
	Purr() string
}

type HouseCat interface {
	Meow() string
}

type TabbyCat struct {
}

type Lion struct {
}

func (tabbyCat *TabbyCat) Purr() string {
	return "*purr*"
}

func (tabbyCat *TabbyCat) Meow() string {
	return "Meow!"
}

func (lion *Lion) Purr() string {
	return "*rumble*"
}
