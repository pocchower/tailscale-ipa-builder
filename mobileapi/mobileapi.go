package mobileapi

import "fmt"

// Greeter is an exported type visible to gomobile bind
type Greeter struct{}

// Hello is an exported method that will be visible to Swift/Obj-C
func (Greeter) Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

