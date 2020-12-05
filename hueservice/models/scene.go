package models

type Scene struct {
	Name   string   `json:"name,omitempty"`
	Type   string   `json:"type,omitempty"`
	Lights []string `json:"lights,omitempty"`
}
