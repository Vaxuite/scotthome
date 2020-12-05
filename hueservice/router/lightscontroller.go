package router

import (
	"gitlab.com/scotthome/hueservice/models"
)

type LightsController interface {
	Rooms() ([]*models.Room, error)

	SetRoomState(groupName string, on bool) error
	SetRoomScene(groupName, sceneName string) error
	SetRoomBri(groupName string, bri uint8) error
}
