package node

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"taverok/hostver/internal/client/app"
	"taverok/hostver/internal/dto"
)

type Service struct {
	AppService app.Service
}

func (it *Service) CurrentNode() (*dto.Node, error) {
	platform, err := dto.CurrentPlatform()
	if err != nil {
		return nil, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	ip, err := it.discoverExternalIP()
	if err != nil {
		return nil, err
	}

	apps, err := it.AppService.ParseApps(platform)
	if err != nil {
		return nil, err
	}

	return &dto.Node{
		Hostname: hostname,
		Platform: *platform,
		IP:       ip,
		Apps:     apps,
	}, nil
}

func (it *Service) discoverExternalIP() (string, error) {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	ip := struct {
		IP string `json:"query"`
	}{}

	err = json.Unmarshal(body, &ip)
	if err != nil {
		return "", err
	}

	return ip.IP, nil
}
