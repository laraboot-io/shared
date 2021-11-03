package integration

import (
	"fmt"
	"github.com/onsi/gomega/format"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testOffline(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
		pack   occam.Pack
		docker occam.Docker
	)
	format.MaxLength = 0
	SetDefaultEventuallyTimeout(10 * time.Second)

	it.Before(func() {
		pack = occam.NewPack()
		docker = occam.NewDocker()

		PreparePhpOfflineBps()
	})

	it.After(func() {
		Expect(os.RemoveAll(sharedOfflineURI)).To(Succeed())
		Expect(os.RemoveAll(phpDistOfflineURI)).To(Succeed())
		Expect(os.RemoveAll(phpWebOfflineURI)).To(Succeed())
	})

	context("when offline", func() {
		var (
			image     occam.Image
			container occam.Container
			name      string
			source    string
		)

		it.Before(func() {
			var err error
			name, err = occam.RandomName()
			Expect(err).NotTo(HaveOccurred())
		})

		it.After(func() {
			Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
			Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
			Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
			Expect(os.RemoveAll(source)).To(Succeed())
		})

		it("creates a working OCI image that serves a simple php application", func() {
			var err error
			source, err = occam.Source(filepath.Join("testdata", "sandbox"))
			Expect(err).NotTo(HaveOccurred())

			var logs fmt.Stringer
			image, logs, err = pack.WithNoColor().Build.
				WithPullPolicy("always").
				WithTrustBuilder().
				WithBuildpacks(phpDistOfflineURI, phpComposerOfflineURI, phpWebOfflineURI).
				WithNetwork("default").
				Execute(name, source)
			Expect(err).NotTo(HaveOccurred(), logs.String())

			Expect(logs.String()).To(ContainSubstring(buildpackInfo.Buildpack.Name))
			Expect(logs.String()).NotTo(ContainSubstring("Downloading"))

			container, err = docker.Container.Run.
				WithEnv(map[string]string{"PORT": "8080"}).
				WithPublish("8080").
				Execute(image.ID)
			Expect(err).ToNot(HaveOccurred())
		})
	})
}
