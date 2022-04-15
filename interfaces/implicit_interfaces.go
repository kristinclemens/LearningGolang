package interfaces

// An example of implicit interfaces. TabbyCat implements Cat and HouseCat implicitly because it has methods satisfying
// the signature of both interfaces. Lion implements only Cat because it does NOT have a method satisfying the HouseCat
// signature. Both TabbyCat and Lion implement Animal because Animal has no functions.
// Ref: https://dev.to/mraszplewicz/golang-through-the-eyes-of-a-java-developer-pros-and-cons-25o2

type Animal interface{}

// Cat can Purr
type Cat interface {
	Purr() string
}

// HouseCat can Meow
type HouseCat interface {
	Meow() string
}

type TabbyCat struct {
}

type Lion struct {
}

func (tabbyCat TabbyCat) Purr() string {
	return "*purr*"
}

func (tabbyCat TabbyCat) Meow() string {
	return "Meow!"
}

func (lion Lion) Purr() string {
	return "*rumble*"
}
