package main

import (
	"os"

	"github.com/laraboot-io/shared"
	"github.com/laraboot-io/shared/cmd/spec"
	"github.com/paketo-buildpacks/packit"
)

func main() {
	logEmitter := shared.NewLogEmitter(os.Stdout)
	packit.Build(laraboot.Build(logEmitter))
}
