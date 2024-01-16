package dto

import "taverok/hostver/pkg/semver"

type Version struct {
	Version string `json:"version"`
	semver.Semver
}

func NewVersion(version string) Version {
	sv, _ := semver.Parse(version)

	return Version{
		Version: version,
		Semver:  sv,
	}
}
