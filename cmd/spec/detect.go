package laraboot

import (
	"github.com/paketo-buildpacks/packit"
)

// Detect .
func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (
		packit.DetectResult,
		error,
	) {
		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: "shared-lib"},
				},
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "shared-lib",
						Metadata: map[string]string{
							"version-source": "laraboot.json",
						},
					},
					{
						Name: "php",
						Metadata: map[string]bool{
							"build":  true,
							"launch": true,
						},
					},
					{
						Name: "composer",
						Metadata: map[string]bool{
							"build":  true,
							"launch": true,
						},
					},
				},
			},
		}, nil
	}
}
