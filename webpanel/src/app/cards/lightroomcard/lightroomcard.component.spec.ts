import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LightRoomCardComponent } from './lightroomcard.component';

describe('LightroomcardComponent', () => {
  let component: LightRoomCardComponent;
  let fixture: ComponentFixture<LightRoomCardComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LightRoomCardComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LightRoomCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
