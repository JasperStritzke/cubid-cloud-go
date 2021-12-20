package template

import (
	"errors"
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"strings"
)

type VersionValue struct {
	BuildURL    string `json:"url"`
	Display     string `json:"display"`
	IsProxyBool bool   `json:"isProxy"`
	Mention     string `json:"mention"`
}

func (v *VersionValue) IsProxy() bool {
	return v.IsProxyBool
}
func (v *VersionValue) IsServer() bool {
	return !v.IsProxyBool
}

func (v *VersionValue) DownloadToPath(dirName string) error {
	return fileutil.DownloadFile(dirName, v.BuildURL)
}

func (v VersionValue) TypeAsString() string {
	var typeString = "Server"
	if v.IsProxyBool {
		typeString = "IsProxyBool"
	}

	return typeString
}

type Versions struct {
	BungeeCord VersionValue
	Waterfall  VersionValue
	Paper17    VersionValue
	Paper16    VersionValue
	Empty      VersionValue
}

var Version = Versions{
	BungeeCord: VersionValue{
		BuildURL:    "https://ci.md-5.net/job/BungeeCord/lastSuccessfulBuild/artifact/bootstrap/target/BungeeCord.jar",
		Display:     "BungeeCord",
		Mention:     "Last build",
		IsProxyBool: true,
	},
	Waterfall: VersionValue{
		BuildURL:    "https://papermc.io/api/v2/projects/waterfall/versions/1.17/builds/451/downloads/waterfall-1.17-454.jar",
		Display:     "Waterfall-1.17",
		Mention:     "build 454 (probably outdated) 18 dec 2021",
		IsProxyBool: true,
	},
	Paper17: VersionValue{
		BuildURL:    "https://papermc.io/api/v2/projects/paper/versions/1.17.1/builds/334/downloads/paper-1.17.1-401.jar",
		Display:     "Paper-1.17",
		Mention:     "build 401 (probably outdated) 18 dec 2021",
		IsProxyBool: false,
	},
	Paper16: VersionValue{
		BuildURL:    "https://papermc.io/api/v2/projects/paper/versions/1.16.5/builds/788/downloads/paper-1.16.5-794.jar",
		Display:     "Paper-1.16",
		Mention:     "build 794 (probably outdated) 18 dec 2021",
		IsProxyBool: false,
	},
}

var VersionsAsArray = []VersionValue{Version.BungeeCord, Version.Waterfall, Version.Paper17, Version.Paper16}

func GetVersionByString(s string) (VersionValue, error) {
	for _, version := range VersionsAsArray {
		if strings.ToLower(version.Display) == s {
			return version, nil
		}
	}

	return VersionValue{}, errors.New("provided version not pre configured. please use createEmpty or specify a valid version")
}
