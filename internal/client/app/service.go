package app

import (
	"encoding/json"
	"strings"
	"taverok/hostver/internal/dto"
	"taverok/hostver/pkg/exec"
)

type Service struct {
}

func (it *Service) ParseApps(p *dto.Platform) ([]dto.App, error) {
	var apps []dto.App
	var err error
	switch p.OS {
	case dto.OSMacOS:
		apps, err = it.parseMacApps()
	case dto.OSLinux:
		switch strings.TrimSpace(strings.ToLower(p.OsName)) {
		case "ubuntu":
			apps, err = it.parseUbuntuApps()
		case "debian":
			apps, err = it.parseUbuntuApps()
		default:
			apps, err = it.parseLinuxApps()
		}
	case dto.OSWindows:
		apps, err = it.parseWinApps()
	}

	return apps, err
}

func (it *Service) parseMacApps() ([]dto.App, error) {
	rr := struct {
		Apps []struct {
			Name    string `json:"_name"`
			Version string `json:"version"`
		} `json:"SPApplicationsDataType"`
	}{}

	raw, err := exec.SafeExec("system_profiler", "SPApplicationsDataType", "-json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(raw), &rr)
	if err != nil {
		return nil, err
	}

	apps := make([]dto.App, 0, len(rr.Apps))
	for _, r := range rr.Apps {
		app := dto.App{
			Name:    r.Name,
			Version: dto.NewVersion(r.Version),
		}
		apps = append(apps, app)
	}

	return apps, nil
}

func (it *Service) parseUbuntuApps() ([]dto.App, error) {
	panic("implement me")
}

func (it *Service) parseLinuxApps() ([]dto.App, error) {
	return it.parseUbuntuApps()
}

func (it *Service) parseWinApps() ([]dto.App, error) {
	panic("implement me")
}
