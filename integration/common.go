package integration

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/cloudfoundry/dagger"
	"github.com/paketo-buildpacks/occam"

	. "github.com/onsi/gomega"
)

var (
	sharedOfflineURI      string
	phpDistOfflineURI     string
	phpWebOfflineURI      string
	phpComposerOfflineURI string
	smokeGunPackage       struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	buildpackInfo struct {
		Buildpack struct {
			ID   string
			Name string
		}
	}
)

func PreparePhpOfflineBps() {
	var config struct {
		PhpDistOffline string `json:"php-dist"`
	}

	file, err := os.Open("../integration.json")
	Expect(err).ToNot(HaveOccurred())
	defer file.Close()
	Expect(json.NewDecoder(file).Decode(&config)).To(Succeed())

	var cmdConfig struct {
		SmokeGunPackage struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"smoke-gun"`
	}

	// test-case
	file, err = os.Open("../integration/testdata/sandbox/shared.json")
	defer file.Close()
	Expect(json.NewDecoder(file).Decode(&cmdConfig)).To(Succeed())
	smokeGunPackage = cmdConfig.SmokeGunPackage

	bpRoot, err := filepath.Abs("./..")
	Expect(err).ToNot(HaveOccurred())

	sharedOfflineURI, err = Package(bpRoot, "1.2.3", true)
	Expect(err).ToNot(HaveOccurred())

	buildpackStore := occam.NewBuildpackStore()

	phpDistOfflineURI, err = buildpackStore.Get.
		WithVersion("1.2.3").
		WithOfflineDependencies().
		Execute(config.PhpDistOffline)

	phpWebRepo, err := dagger.GetLatestUnpackagedCommunityBuildpack("paketo-buildpacks", "php-web")
	Expect(err).NotTo(HaveOccurred())

	phpWebOfflineURI, err = Package(phpWebRepo, "1.2.3", true)
	Expect(err).ToNot(HaveOccurred())

	phpComposerRepo, err := dagger.GetLatestUnpackagedCommunityBuildpack("paketo-buildpacks", "php-composer")
	Expect(err).NotTo(HaveOccurred())

	phpComposerOfflineURI, err = Package(phpComposerRepo, "1.2.3", true)
	Expect(err).ToNot(HaveOccurred())
}

// PreparePhpBps builds the current buildpacks.
func PreparePhpBps() ([]string, error) {
	var config struct {
		PhpDist string `json:"php-dist"`
	}

	file, err := os.Open("../integration.json")
	Expect(err).ToNot(HaveOccurred())
	defer file.Close()

	Expect(json.NewDecoder(file).Decode(&config)).To(Succeed())

	bpRoot, err := filepath.Abs("./..")
	if err != nil {
		return []string{}, err
	}

	composerBp, err := Package(bpRoot, "1.2.3", false)
	if err != nil {
		return []string{}, err
	}

	buildpackStore := occam.NewBuildpackStore()
	phpDistBp, err := buildpackStore.Get.Execute(config.PhpDist)
	Expect(err).ToNot(HaveOccurred())

	phpWebBp, err := dagger.GetLatestBuildpack("php-web-cnb")
	Expect(err).NotTo(HaveOccurred())

	return []string{phpDistBp, composerBp, phpWebBp}, nil
}

// MakeBuildEnv creates a build environment map
func MakeBuildEnv(debug bool) map[string]string {
	env := make(map[string]string)
	if debug {
		env["BP_DEBUG"] = "true"
	}
	githubToken := os.Getenv("GIT_TOKEN")
	if githubToken != "" {
		env["COMPOSER_GITHUB_OAUTH_TOKEN"] = githubToken
	}

	return env
}

func DecodeBPToml() {
	file, err := os.Open("../buildpack.toml")
	Expect(err).NotTo(HaveOccurred())
	defer file.Close()

	_, err = toml.DecodeReader(file, &buildpackInfo)
	Expect(err).NotTo(HaveOccurred())
}

func Package(root, version string, cached bool) (string, error) {
	var cmd *exec.Cmd

	dir, err := filepath.Abs("./..")
	if err != nil {
		return "", err
	}

	bpPath := filepath.Join(root, "artifact")
	if cached {
		cmd = exec.Command(filepath.Join(dir, ".bin", "packager"), "--archive", "--version", version, fmt.Sprintf("%s-cached", bpPath))
	} else {
		cmd = exec.Command(filepath.Join(dir, ".bin", "packager"), "--archive", "--uncached", "--version", version, bpPath)
	}

	cmd.Env = append(os.Environ(), fmt.Sprintf("PACKAGE_DIR=%s", bpPath))
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	if cached {
		return fmt.Sprintf("%s-cached.tgz", bpPath), nil
	}

	return fmt.Sprintf("%s.tgz", bpPath), nil
}
