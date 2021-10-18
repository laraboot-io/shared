package shared

import (
	"bytes"
	"fmt"
	. "github.com/onsi/gomega"
	"github.com/paketo-buildpacks/packit"
	"github.com/sclevine/spec"
	"testing"
)

func TestLogEmitter(
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
				buffer  *bytes.Buffer
				emitter LogEmitter
			)
			it.Before(func() {
				buffer = bytes.NewBuffer(nil)
				emitter = NewLogEmitter(buffer)
			})

			it("prints details about the environment", func() {
				emitter.Environment(packit.Environment{
					"GEM_PATH.override": "/some/path",
				})
				Expect(buffer.String()).To(ContainSubstring("    GEM_PATH -> \"/some/path\""))
			})
		},
	)
}

func fail(message string, skip ...int) {
	fmt.Println(message)
	fmt.Println(skip)
}
