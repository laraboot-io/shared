package shared

import (
	"io/ioutil"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/packit"
	"github.com/sclevine/spec"
)

func TestContributor(t *testing.T) {
	t.Helper()
	RegisterFailHandler(fail)
	spec.Run(t, "RunContributor",
		func(
			t *testing.T,
			when spec.G,
			it spec.S) {
			var (
				err          error
				Expect       = NewWithT(t).Expect
				layersDir    string
				cnbDir       string
				somePackage  Package
				layer        packit.Layer
				buildContext packit.BuildContext
			)
			it.Before(func() {
				layersDir, err = ioutil.TempDir("", "layers")
				Expect(err).NotTo(HaveOccurred())
				cnbDir, err = ioutil.TempDir("", "cnb")
				Expect(err).NotTo(HaveOccurred())
				buildContext = packit.BuildContext{
					CNBPath: cnbDir,
					Stack:   "some-stack",
					BuildpackInfo: packit.BuildpackInfo{
						Name:    "Some Shared Buildpack",
						Version: "1.2.3",
					},
					Plan: packit.BuildpackPlan{
						Entries: []packit.BuildpackPlanEntry{
							{
								Name:     "shared",
								Metadata: map[string]interface{}{},
							},
						},
					},
					Platform: packit.Platform{Path: "platform"},
					Layers:   packit.Layers{Path: layersDir},
				}
				layer = packit.Layer{
					Path:   cnbDir,
					Name:   "myLayer",
					Build:  false,
					Launch: false,
					Cache:  false,
				}
			})
			it("succeeds", func() {
				somePackage, err = NewPackage("someorg/somepackage", buildContext, layer)
				Expect(err).To(BeNil())
				Expect(somePackage).NotTo(BeNil())
				install, err := somePackage.Install()
				Expect(err).To(BeNil())
				Expect(string(install)).To(ContainSubstring("ok"))
			})
		},
	)
}