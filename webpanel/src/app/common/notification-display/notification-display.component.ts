import { Component, Inject, Input, OnInit } from '@angular/core';
import { MAT_SNACK_BAR_DATA } from '@angular/material/snack-bar';

@Component({
  selector: 'app-notification-display',
  templateUrl: './notification-display.component.html',
  styleUrls: ['./notification-display.component.less']
})
export class NotificationDisplayComponent  {

  constructor(@Inject(MAT_SNACK_BAR_DATA) public data: any) {}

}
