package mobileapi

import "fmt"

// Hello is a basic example function exported to iOS.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

