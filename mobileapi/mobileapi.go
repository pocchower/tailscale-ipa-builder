package mobileapi

type Greeter struct{}

func (Greeter) Hello(name string) string {
	return "Hello, " + name
}

