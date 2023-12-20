package shared

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/cloudfoundry/libcfbuildpack/helper"
	"github.com/paketo-buildpacks/packit/v2"
)

// Contributor represents a PHP contribution by the buildpack.
type Contributor struct {
	name          string
	version       string
	global        bool
	logger        LogEmitter
	customIniPath string
	context       packit.BuildContext
	layer         packit.Layer
	layers        packit.Layers
}

// NewGlobalContributor creates a Contributor instance and set global = true.
func NewGlobalContributor(name string, context packit.BuildContext, layer packit.Layer) (Contributor, error) {
	newPackage, err := NewContributor(name, context, layer)
	if err != nil {
		return Contributor{}, err
	}
	newPackage.global = true
	return newPackage, nil
}

// NewContributor creates a new Contributor instance.
func NewContributor(name string, context packit.BuildContext, layer packit.Layer) (Contributor, error) {
	packageName := name
	version := "latest"

	if strings.Contains(name, ":") {
		tokens := strings.Split(name, ":")
		packageName = tokens[0]
		version = tokens[1]
	}

	contributor := Contributor{
		name:          packageName,
		version:       version,
		layer:         layer,
		layers:        context.Layers,
		logger:        NewLogEmitter(os.Stdout),
		customIniPath: fmt.Sprintf("%s/php.ini", layer.Path),
		context:       context,
	}

	return contributor, nil
}

// WriteCustomInitFile writes an ini file.
func (l Contributor) WriteCustomInitFile(templateBody string, outputPath string, data interface{}) error {
	t, err := template.New(filepath.Base(outputPath)).Parse(templateBody)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	err = t.Execute(&b, data)
	if err != nil {
		return err
	}
	return helper.WriteFileFromReader(outputPath, 0o644, &b) //nolint:gomnd //ignore
}

// Install the package
// It requieres `composer buildpack` to run the installation.
func (l Contributor) Install() ([]byte, error) {
	l.logger.Detail("Installing %s (%s):", l.name, l.version)
	composerLayerPath := "/layers/paketo-buildpacks_php-composer/composer"
	composerPath := fmt.Sprintf("%s/composer.phar", composerLayerPath)
	err := l.WriteCustomInitFile(`extension=openssl
extension=mbstring
extension=fileinfo
extension=curl`,
		l.customIniPath,
		"")
	if err != nil {
		return nil, err
	}

	// installation args
	args := []string{
		fmt.Sprintf("-dextension_dir=%s", os.Getenv("PHP_EXTENSION_DIR")),
		fmt.Sprintf("-derror_reporting=%s", "E_ALL"),
		"-c",
		l.customIniPath,
		composerPath,
	}

	if l.global {
		args = append(args, []string{"global"}...)
	}

	args = append(args, []string{
		"require",
		fmt.Sprintf("%s %s", l.name, l.version),
		"--prefer-stable",
		"-W",
	}...)

	command, err := RunCommand(l.context, "php", args...)
	if err != nil {
		return command, err
	}

	l.logger.Break()
	return command, nil
}
