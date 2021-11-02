package main

import (
	"os"

	"github.com/laraboot-io/shared"
	"github.com/laraboot-io/shared/cmd/spec"
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/chronos"
)

func main() {
	logEmitter := shared.NewLogEmitter(os.Stdout)
	packit.Build(laraboot.Build(
		logEmitter,
		chronos.DefaultClock))
}
