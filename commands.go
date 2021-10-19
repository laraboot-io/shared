package shared

import (
	"github.com/paketo-buildpacks/packit"
	"os/exec"
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
