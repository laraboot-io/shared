package shared

import (
	"fmt"
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/pexec"
	"os"
	"os/exec"
	"strings" //nolint:gci // .
)

// GitCmd runs git command.
func GitCmd(args ...string) error {
	git := pexec.NewExecutable("git")

	fmt.Printf("Running git command : `git %s` \n", strings.Join(args, " "))

	gitErr := git.Execute(pexec.Execution{
		Args:   args,
		Stdout: os.Stdout,
	})

	if gitErr != nil {
		return gitErr
	}

	return nil
}

// ComposerCommand runs a composer command .
func ComposerCommand(context packit.BuildContext, customIni string, arg ...string) ([]byte, error) {
	phpPath, err := exec.LookPath("php")
	composerDir := "/layers/paketo-buildpacks_php-composer/composer"
	composerPhar := fmt.Sprintf("%s/composer.phar", composerDir)

	if err != nil {
		panic(err)
	}

	args := append([]string{
		fmt.Sprintf("-dextension_dir=%s", os.Getenv("PHP_EXTENSION_DIR")),
		fmt.Sprintf("-derror_reporting=%s", "E_ALL"),
		"-c",
		customIni,
		composerPhar,
	}, arg...)

	cmd := exec.Command(phpPath, args...) //nolint:gosec //ignore
	cmd.Dir = context.WorkingDir

	return cmd.CombinedOutput()
}
