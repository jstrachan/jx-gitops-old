package plugins

import (
	"fmt"
	"strings"

	jenkinsv1 "github.com/jenkins-x/jx/v2/pkg/apis/jenkins.io/v1"
	"github.com/jenkins-x/jx/v2/pkg/extensions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Platform represents a platform for binaries
type Platform struct {
	Goarch string
	Goos   string
}

const (
	// HelmPluginName the default name of the helm plugin
	HelmPluginName = "helm"
)

var (
	defaultPlatforms = []Platform{
		{
			Goarch: "amd64",
			Goos:   "Windows",
		},
		{
			Goarch: "amd64",
			Goos:   "Darwin",
		},
		{
			Goarch: "amd64",
			Goos:   "Linux",
		},
		{
			Goarch: "arm",
			Goos:   "Linux",
		},
		{
			Goarch: "386",
			Goos:   "Linux",
		},
	}
)

// Extension returns the default distribution extension; `tar.gz` or `zip` for windows
func (p Platform) Extension() string {
	if p.IsWindows() {
		return "zip"
	}
	return "tar.gz"
}

// IsWindows returns true if the platform is windows
func (p Platform) IsWindows() bool {
	return p.Goos == "Windows"
}

// GetHelmBinary returns the path to the locally installed helm 3 extension
func GetHelmBinary(version string) (string, error) {
	if version == "" {
		version = HelmVersion
	}
	plugin := CreateHelmPlugin(version)
	return extensions.EnsurePluginInstalled(plugin)
}

// CreateHelmPlugin creates the helm 3 plugin
func CreateHelmPlugin(version string) jenkinsv1.Plugin {
	binaries := CreateBinaries(func(p Platform) string {
		return fmt.Sprintf("https://get.helm.sh/helm-v%s-%s-%s.%s", version, strings.ToLower(p.Goos), strings.ToLower(p.Goarch), p.Extension())
	})

	plugin := jenkinsv1.Plugin{
		ObjectMeta: metav1.ObjectMeta{
			Name: HelmPluginName,
		},
		Spec: jenkinsv1.PluginSpec{
			SubCommand:  "helm",
			Binaries:    binaries,
			Description: "helm 3 binary",
			Name:        HelmPluginName,
			Version:     version,
		},
	}
	return plugin
}

// CreateBinaries a helper function to create the binary resources for the platforms for a given callback
func CreateBinaries(createURLFn func(Platform) string) []jenkinsv1.Binary {
	var answer []jenkinsv1.Binary
	for _, p := range defaultPlatforms {
		u := createURLFn(p)
		if u != "" {
			answer = append(answer, jenkinsv1.Binary{
				Goarch: p.Goarch,
				Goos:   p.Goos,
				URL:    u,
			})
		}
	}
	return answer
}
