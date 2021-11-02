// Package laraboot ...
package laraboot

import (
	"github.com/laraboot-io/shared"
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/chronos"
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
func Build(logger shared.LogEmitter, clock chronos.Clock) packit.BuildFunc { //nolint:funlen // .
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)
		layer, err := context.Layers.Get("shared-lib")
		if err != nil {
			return packit.BuildResult{}, err
		}

		newPackage, err := shared.NewPackage("monolog", context, layer)

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
