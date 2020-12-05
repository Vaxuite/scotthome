package models

type GroupModification struct {
	GroupName  string `json:"groupName"`
	On         bool   `json:"on"`
	SceneName  string `json:"sceneName"`
	Brightness uint8  `json:"bri"`
}
