import {Component} from '@angular/core';
import {BreakpointObserver} from '@angular/cdk/layout';
import {NotificationService} from './common/service/notification.service';
import {INotification} from './common/domain/errors';
import {MatSnackBar} from '@angular/material/snack-bar';
import {NotificationDisplayComponent} from './common/notification-display/notification-display.component';
import {Router} from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.less']
})
export class AppComponent {


  constructor(public notificationService: NotificationService,
              private snackBar: MatSnackBar,
              private router: Router) {
    notificationService.notificationRaised$.subscribe(notification => {
      this.raiseNotification(notification);
    });
  }

  public raiseNotification(notification: INotification) {
    this.snackBar.openFromComponent(NotificationDisplayComponent, {
      duration: 3000,
      data: {
        message: notification.message,
        type: notification.type,
        duration: 3000
      }
    });
  }
}
