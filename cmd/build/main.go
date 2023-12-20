package main

import (
	"os"

	"github.com/laraboot-io/shared"
	laraboot "github.com/laraboot-io/shared/cmd/spec"
	packit "github.com/paketo-buildpacks/packit/v2"
)

func main() {
	logEmitter := shared.NewLogEmitter(os.Stdout)
	packit.Build(laraboot.Build(logEmitter))
}
