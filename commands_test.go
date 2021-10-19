package shared

import (
	"fmt"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/packit"
	"github.com/sclevine/spec"
	"io/ioutil"
	"testing"
)

func TestCommands(
	t *testing.T,
) {
	t.Helper()
	RegisterFailHandler(fail)
	spec.Run(
		t,
		"Run",
		func(
			t *testing.T,
			when spec.G,
			it spec.S) {

			var (
				Expect = NewWithT(t).Expect

				layersDir    string
				cnbDir       string
				buildContext packit.BuildContext
			)

			it.Before(func() {
				var err error
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
			})

			it("runs a command and returns the output", func() {
				_, err := RunCommand(buildContext, "ls")
				Expect(err).To(BeNil())
				if err != nil {
					_ = fmt.Errorf("[ERROR] %s", err)
				}
			})
		},
	)
}

func fail(message string, skip ...int) {
	fmt.Println(message)
	fmt.Println(skip)
}
