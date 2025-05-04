// Trigger: dummy change to re-run GitHub Actions
cd ~/tailscale-ipa-builder/mobileapi
nano mobileapi.go
package mobileapi

type Greeter struct{}

func (Greeter) Hello(name string) string {
	return "Hello, " + name
}

