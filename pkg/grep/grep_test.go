package grep

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchFirst_sample1(t *testing.T) {
	sample := `Software:

    System Software Overview:

      System Version: macOS 14.2.1 (23C71)
      Kernel Version: Darwin 23.2.0
      Boot Volume: root
      Boot Mode: Normal
      Computer Name: AAA’s MacBook Pro
      User Name: AAA BBB (ABAB)
      Secure Virtual Memory: Enabled
      System Integrity Protection: Enabled
      Time since boot: 11 days, 2 hours, 30 minutes`

	cases := []struct {
		pattern string
		want    []string
	}{
		{"Soft(.+)", []string{"ware:"}},
		{"System(.+)", []string{" Software Overview:"}},
		{"Kernel (.+):", []string{"Version"}},
		{"Kernel Version: Darwin ([0-9.]+)", []string{"23.2.0"}},
	}

	for _, c := range cases {
		out := MatchFirst(sample, regexp.MustCompile(c.pattern))
		assert.Equal(t, c.want, out)
	}
}
