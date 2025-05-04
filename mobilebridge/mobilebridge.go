package mobilebridge

import (
    "fmt"
)

// Controller is the Swift-callable VPN manager
type Controller struct{}

// NewController returns a new instance
func NewController() *Controller {
    return &Controller{}
}

// Start simulates VPN startup
func (c *Controller) Start() string {
    return "Tailscale backend started (stub)"
}

