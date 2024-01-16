package dto

import "encoding/json"

type Node struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Platform
	Apps []App `json:"apps"`
}

func (it *Node) AsJson() (string, error) {
	indent, err := json.MarshalIndent(*it, "", "  ")
	if err != nil {
		return "", err
	}

	return string(indent), nil
}
