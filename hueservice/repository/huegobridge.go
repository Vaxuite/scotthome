package repository

import "github.com/amimof/huego"

type HueGoBridge interface {
	GetLights() ([]huego.Light, error)
	GetGroups() ([]huego.Group, error)
	GetScenes() ([]huego.Scene, error)
}
