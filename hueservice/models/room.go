package models

type Room struct {
	Name      string     `json:"name,omitempty"`
	Lights    []string   `json:"lights,omitempty"`
	Type      string     `json:"type,omitempty"`
	Scenes    []*Scene   `json:"scenes,omitempty"`
	ID        int        `json:"-"`
	RoomState *RoomState `json:"roomState"`
}
