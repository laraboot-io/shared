package main

import (
	laraboot "github.com/laraboot-io/shared/cmd/spec"
	"github.com/paketo-buildpacks/packit"
)

func main() {
	packit.Detect(laraboot.Detect())
}
