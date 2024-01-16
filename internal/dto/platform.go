package dto

import (
	"fmt"
	"regexp"
	"runtime"
	"taverok/hostver/pkg/exec"
	"taverok/hostver/pkg/grep"
)

type OS string

const (
	OSLinux   = "linux"
	OSWindows = "windows"
	OSMacOS   = "macos"
)

type Platform struct {
	OS            string  `json:"os"`
	OsName        string  `json:"osName"`
	OsVersion     Version `json:"osVersion"`
	KernelVersion Version `json:"kernelVersion"`
	Arch          string  `json:"arch"`
}

func CurrentPlatform() (*Platform, error) {
	var err error

	p := &Platform{
		Arch: runtime.GOARCH,
	}

	// go tool dist list
	switch runtime.GOOS {
	case "windows":
		p.OS = OSWindows
		err = p.parseWindows()
	case "darwin":
		p.OS = OSMacOS
		err = p.parseMac()
	case "linux", "openbsd", "freebsd", "netbsd":
		p.OS = OSLinux
		err = p.parseLinux()
	default:
		panic(fmt.Sprintf("unsupported os: %s", runtime.GOOS))
	}

	return p, err
}

func (it *Platform) parseMac() error {
	it.OsName = "macOS"

	profile, err := exec.SafeExec("system_profiler", "SPSoftwareDataType")
	if err != nil {
		return err
	}

	raw := grep.ConcatFirst(profile, regexp.MustCompile(".+ System Version: macOS ([0-9.]+)"), "")
	it.OsVersion = NewVersion(raw)

	raw = grep.ConcatFirst(profile, regexp.MustCompile("Kernel Version: Darwin ([0-9.]+)"), "")
	it.KernelVersion = NewVersion(raw)

	return nil
}

func (it *Platform) parseWindows() error {
	panic("implement me")
}

func (it *Platform) parseLinux() error {
	panic("implement me")
}
