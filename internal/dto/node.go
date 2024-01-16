package dto

type Node struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Platform
	Apps []App `json:"apps"`
}
