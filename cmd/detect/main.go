package main

import (
	laraboot "github.com/laraboot-io/shared/cmd/spec"
	packit "github.com/paketo-buildpacks/packit/v2"
)

func main() {
	packit.Detect(laraboot.Detect())
}
