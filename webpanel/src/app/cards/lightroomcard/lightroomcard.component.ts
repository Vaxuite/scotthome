import {Component, Input, OnInit} from '@angular/core';
import {LightRoom} from '../../common/domain/lights/room';
import {LightService} from '../../common/service/lightservice.service';

@Component({
  selector: 'app-lightroomcard',
  templateUrl: './lightroomcard.component.html',
  styleUrls: ['./lightroomcard.component.css']
})
export class LightRoomCardComponent implements OnInit {

  @Input() lightRoom: LightRoom;

  constructor(private lightService: LightService) {
  }

  ngOnInit() {
  }

  private setState(on: boolean) {
    this.lightService.setState({bri: 0, sceneName: '', groupName: this.lightRoom.name, on}).subscribe(results => {});
  }

  private setBrightness(bri: number) {
    this.lightService.setBrightness({on: this.lightRoom.roomState.on, bri, sceneName: '',
      groupName: this.lightRoom.name}).subscribe(results => {});
  }

  private setScene(scene: string) {
    this.lightService.setScene({on: this.lightRoom.roomState.on, bri: this.lightRoom.roomState.brightness, sceneName: scene,
      groupName: this.lightRoom.name}).subscribe(results => {});
  }

}
