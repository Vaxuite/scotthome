import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NotificationDisplayComponent } from './notification-display.component';
import { MatIconModule } from '@angular/material/icon';
import { MAT_SNACK_BAR_DATA, MatSnackBarModule, MatSnackBarRef } from '@angular/material/snack-bar';

describe('NotificationDisplayComponent', () => {
  let component: NotificationDisplayComponent;
  let fixture: ComponentFixture<NotificationDisplayComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NotificationDisplayComponent ],
      imports: [
        MatIconModule,
        MatSnackBarModule
      ],
      providers: [{
        provide: MatSnackBarRef,
        useValue: {}
      }, {
        provide: MAT_SNACK_BAR_DATA,
        useValue: {}
      }]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NotificationDisplayComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
