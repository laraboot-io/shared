package shared

import (
	"os/exec"

	"github.com/paketo-buildpacks/packit"
)

// RunCommand runs a composer command .
func RunCommand(context packit.BuildContext, execName string, args ...string) ([]byte, error) {
	path, err := exec.LookPath(execName)

	if err != nil {
		panic(err)
	}

	cmd := exec.Command(path, args...) //nolint:gosec //.
	cmd.Dir = context.WorkingDir

	return cmd.CombinedOutput()
}
