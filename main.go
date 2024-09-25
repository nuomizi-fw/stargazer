//go:generate bash -c "go generate ./ent"
package main

import "github.com/nuomizi-fw/stargazer/cmd"

func main() {
	cmd.Execute()
}
