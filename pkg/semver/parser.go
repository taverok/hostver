package semver

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`(\d+)\.(\d+)\.?(\d+)?[+-]?([a-zA-Z0-9.+-]*)`)

var (
	ErrInvalidVersion = fmt.Errorf("invalid version")
)

type Semver struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

func Parse(ver string) (Semver, error) {
	var majV, minV, patch int

	v64, err := strconv.ParseInt(ver, 10, 64)
	if err == nil {
		return Semver{Major: int(v64)}, nil
	}

	ver = strings.TrimSpace(ver)
	allGroups := pattern.FindAllStringSubmatch(ver, -1)

	if len(allGroups) == 0 || len(allGroups[0]) == 0 {
		return Semver{}, ErrInvalidVersion
	}

	matches := allGroups[0]

	majV, err = strconv.Atoi(matches[1])
	if err != nil {
		return Semver{}, ErrInvalidVersion
	}
	minV, _ = strconv.Atoi(matches[2])
	patch, _ = strconv.Atoi(matches[3])

	return Semver{Major: majV, Minor: minV, Patch: patch}, nil
}
