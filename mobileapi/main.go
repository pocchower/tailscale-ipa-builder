package mobileapi

import "fmt"

// Hello is an exported function accessible from iOS.
func Hello(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

