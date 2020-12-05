import { EventEmitter, Injectable } from '@angular/core';
import { INotification } from '../domain/errors';


@Injectable({
  providedIn: 'root'
})
export class NotificationService {

  public notificationRaised$: EventEmitter<INotification>;

  constructor() {
    this.notificationRaised$ = new EventEmitter();
  }

  public raiseNotification(notification: INotification) {
    this.notificationRaised$.emit(notification);
  }
}
