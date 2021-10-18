package shared

import (
	"bytes"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/packit"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
	"testing"
)

func testLogEmitter(t *testing.T, context spec.G, it spec.S) {

	var (
		Expect = NewWithT(t).Expect

		buffer  *bytes.Buffer
		emitter LogEmitter
	)

	it.Before(func() {
		buffer = bytes.NewBuffer(nil)
		emitter = NewLogEmitter(buffer)
	})

	context("Environment", func() {
		it("prints details about the environment", func() {
			emitter.Environment(packit.Environment{
				"GEM_PATH.override": "/some/path",
			})

			Expect(buffer.String()).To(ContainSubstring("    GEM_PATH -> \"/some/path\""))
		})
	})
}

func TestUnitMRI(t *testing.T) {
	suite := spec.New("shared", spec.Report(report.Terminal{}))
	suite("LogEmitter", testLogEmitter)
	suite.Run(t)
}
