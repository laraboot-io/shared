// Package laraboot ...
package laraboot

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/laraboot-io/shared"
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/postal"
)

type (
	// EntryResolver .
	EntryResolver interface {
		Resolve(string, []packit.BuildpackPlanEntry, []interface{}) (packit.BuildpackPlanEntry, []packit.BuildpackPlanEntry)
		MergeLayerTypes(string, []packit.BuildpackPlanEntry) (launch, build bool)
	}
)

type (
	// DependencyService .
	DependencyService interface {
		Resolve(path, name, version, stack string) (postal.Dependency, error)
		Install(dependency postal.Dependency, cnbPath, layerPath string) error
	}
)

// Build .
func Build(logger shared.LogEmitter) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)
		layer, err := context.Layers.Get("shared-lib")
		if err != nil {
			return packit.BuildResult{}, err
		}

		var config struct {
			SmokeGunPackage struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"smoke-gun"`
		}

		var file, _ = os.Open("./shared.json")
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		if err = json.NewDecoder(file).Decode(&config); err != nil {
			return packit.BuildResult{}, err
		}
		fqq := fmt.Sprintf("%s:%s",
			config.SmokeGunPackage.Name,
			config.SmokeGunPackage.Version)
		newPackage, err := shared.NewPackage(fqq, context, layer)

		if err != nil {
			return packit.BuildResult{}, err
		}

		install, err := newPackage.Install()
		if err != nil {
			logger.Detail(string(install))
			return packit.BuildResult{}, err
		}

		return packit.BuildResult{
			Layers: []packit.Layer{layer},
		}, nil
	}
}
