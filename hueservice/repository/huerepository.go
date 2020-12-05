package repository

import (
	"errors"
	"github.com/amimof/huego"
	"github.com/spf13/viper"
	"strconv"
)

var ErrUnknownGroup = errors.New("unknown group")
var ErrSceneNotFound = errors.New("unknown scene")

type HueGoRepository struct {
	user   string
	bridge HueGoBridge
}

func NewHueGoRepository() (*HueGoRepository, error) {
	hueGo := &HueGoRepository{}
	bridge, _ := huego.Discover()
	bridge = bridge.Login(viper.GetString("hue"))
	hueGo.bridge = bridge
	return hueGo, nil
}

func (hueGoRepository *HueGoRepository) Lights() ([]huego.Light, error) {
	return hueGoRepository.bridge.GetLights()
}

func (hueGoRepository *HueGoRepository) Groups() ([]huego.Group, error) {
	return hueGoRepository.bridge.GetGroups()
}

func (hueGoRepository *HueGoRepository) Group(name string) (*huego.Group, error) {
	groups, err := hueGoRepository.Groups()
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		if group.Name == name {
			return &group, err
		}
	}
	return nil, ErrUnknownGroup
}

func (hueGoRepository *HueGoRepository) Scenes() ([]huego.Scene, error) {
	return hueGoRepository.bridge.GetScenes()
}

func (hueGoRepository *HueGoRepository) Scene(groupName, sceneName string) (*huego.Scene, error) {
	scenes, err := hueGoRepository.bridge.GetScenes()
	if err != nil {
		return nil, err
	}
	group, err := hueGoRepository.Group(groupName)
	if err != nil {
		return nil, err
	}
	for _, scene := range scenes {
		if scene.Group == strconv.Itoa(group.ID) && scene.Name == sceneName {
			return &scene, nil
		}
	}
	return nil, ErrSceneNotFound
}

func (hueGoRepository *HueGoRepository) SetGroupState(groupName string, on bool) error {
	group, err := hueGoRepository.Group(groupName)
	if err != nil {
		return err
	}
	if on {
		return group.On()
	} else {
		return group.Off()
	}
}

func (hueGoRepository *HueGoRepository) SetGroupScene(groupName, sceneName string) error {
	group, err := hueGoRepository.Group(groupName)
	if err != nil {
		return err
	}
	scene, err := hueGoRepository.Scene(groupName, sceneName)
	if err != nil {
		return err
	}
	return group.Scene(scene.ID)
}

func (hueGoRepository *HueGoRepository) SetGroupBri(groupName string, bri uint8) error {
	group, err := hueGoRepository.Group(groupName)
	if err != nil {
		return err
	}
	return group.Bri(bri)
}
