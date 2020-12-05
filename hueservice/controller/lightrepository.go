package controller

import "github.com/amimof/huego"

type LightRepository interface {
	Lights() ([]huego.Light, error)
	Groups() ([]huego.Group, error)
	Scenes() ([]huego.Scene, error)

	SetGroupScene(groupName, sceneName string) error
	SetGroupBri(groupName string, bri uint8) error
	SetGroupState(groupName string, on bool) error
}
