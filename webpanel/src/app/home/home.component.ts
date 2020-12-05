import { Component, OnInit } from '@angular/core';
import {LightService} from '../common/service/lightservice.service';
import {LightRoom} from '../common/domain/lights/room';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent  {

  public lightRooms: LightRoom[] = [];

  constructor(private lightService: LightService) {
   lightService.getRooms().subscribe(results => {
     results.sort((a, b) => a.name.localeCompare(b.name));
     this.lightRooms = results;
     console.log(this.lightRooms);
   });
  }

}
