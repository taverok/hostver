package semver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	cases := []struct {
		version  string
		expected Semver
		err      error
	}{
		{"1.2.3", Semver{Major: 1, Minor: 2, Patch: 3}, nil},
		{"1.2.3-alpha", Semver{Major: 1, Minor: 2, Patch: 3}, nil},
		{"1.2.3-alpha+build", Semver{Major: 1, Minor: 2, Patch: 3}, nil},
		{"illegal:version", Semver{}, ErrInvalidVersion},
		{"245.2", Semver{Major: 245, Minor: 2}, nil},
		{"245", Semver{Major: 245}, nil},

		// semver in text
		{"docker version 1.2.3", Semver{Major: 1, Minor: 2, Patch: 3}, nil},
		{"Docker version 1.2.3, build 100c701", Semver{Major: 1, Minor: 2, Patch: 3}, nil},
	}

	for _, c := range cases {
		t.Run(c.version, func(t *testing.T) {
			actual, err := Parse(c.version)
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.expected, actual)
		})
	}
}
