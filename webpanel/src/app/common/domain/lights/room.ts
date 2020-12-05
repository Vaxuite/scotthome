export interface LightRoom {
  name: string;
  lights: number[];
  roomState: RoomState;
  scenes: Scene[];
}

export interface Scene {
  name: string;
  group: string;
}

export interface RoomState {
  on: boolean;
  brightness: number;
}

export interface GroupModification {
  groupName: string;
  on: boolean;
  sceneName: string;
  bri: number;
}
