package mobileapi

type Greeter struct{}

// Hello is an exported method accessible by gomobile bind
func (Greeter) Hello(name string) string {
	return "Hello, " + name
}

