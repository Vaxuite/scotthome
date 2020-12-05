package controller

import (
	"github.com/amimof/huego"
	"gitlab.com/scotthome/hueservice/models"
	"gitlab.com/scotthome/hueservice/repository"
	"strconv"
)

type DefaultLightsController struct {
	lightRepository LightRepository
}

func NewDefaultLightsController() (*DefaultLightsController, error) {
	hueGo := &DefaultLightsController{}
	lightRepo, err := repository.NewHueGoRepository()
	if err != nil {
		return nil, err
	}
	hueGo.lightRepository = lightRepo
	return hueGo, nil
}

func (defaultLightsController *DefaultLightsController) Lights() ([]huego.Light, error) {
	return defaultLightsController.lightRepository.Lights()
}

func (defaultLightsController *DefaultLightsController) Rooms() ([]*models.Room, error) {
	groups, err := defaultLightsController.lightRepository.Groups()
	if err != nil {
		return nil, err
	}
	rooms := make([]*models.Room, 0)
	for _, group := range groups {
		if group.Type == "Room" {
			rooms = append(rooms, &models.Room{ID: group.ID, Name: group.Name, Lights: group.Lights, Type: group.Type,
				RoomState: &models.RoomState{On: group.GroupState.AllOn, Brightness: group.State.Bri}})
		}
	}
	scenes, err := defaultLightsController.lightRepository.Scenes()
	for _, scene := range scenes {
		for _, room := range rooms {
			if strconv.Itoa(room.ID) == scene.Group {
				room.Scenes = append(room.Scenes, &models.Scene{
					Name:   scene.Name,
					Type:   scene.Type,
					Lights: scene.Lights,
				})
			}
		}
	}
	return rooms, err
}

func (defaultLightsController *DefaultLightsController) SetRoomState(groupName string, on bool) error {
	return defaultLightsController.lightRepository.SetGroupState(groupName, on)
}

func (defaultLightsController *DefaultLightsController) SetRoomScene(groupName, sceneName string) error {
	return defaultLightsController.lightRepository.SetGroupScene(groupName, sceneName)
}

func (defaultLightsController *DefaultLightsController) SetRoomBri(groupName string, bri uint8) error {
	return defaultLightsController.lightRepository.SetGroupBri(groupName, bri)
}
