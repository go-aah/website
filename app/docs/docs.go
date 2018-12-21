package docs

import (
	"path/filepath"

	"aahframe.work"
	"aahframe.work/essentials"
)

const (
	virtualBaseDir = "/aah/documentation"
)

var (
	releases []string
)

// LatestRelease method returns the aah docs latest release number.
func LatestRelease() string {
	return releases[0]
}

// LatestRelease method returns true if given version is latest
// otherwise false.
func IsLatestRelease(v string) bool {
	return LatestRelease() == v
}

// Releases method returns the all the docs release version from config.
func Releases() []string {
	return releases
}

// ReleaseExists method returns true if given docs version exists in the
// release config otherwise false.
func ReleaseExists(v string) bool {
	return ess.IsSliceContainsString(releases, v)
}

// BaseDir method returns the aah documentation physical based directory.
func BaseDir() string {
	return filepath.Join(aah.App().Config().StringDefault("docs.dir", ""), "aah-documentation")
}

// VirutalBaseDir method returns the docs virtual base dir.
func VirutalBaseDir() string {
	return virtualBaseDir
}

// VersionBaseDir method returns the documentation dir path for given
// language and version.
func VersionBaseDir(v string) string {
	return filepath.Join(BaseDir(), v)
}

// LoadFromConfig method loads required value from configuration and others
func LoadFromConfig(_ *aah.Event) {
	app := aah.App()
	releases, _ = app.Config().StringList("docs.releases")

	physicalBaseDir := BaseDir()
	_ = ess.MkDirAll(physicalBaseDir, 0755)

	if err := app.VFS().AddMount(VirutalBaseDir(), physicalBaseDir); err != nil {
		app.Log().Fatal("vfs:", err)
	}
}
