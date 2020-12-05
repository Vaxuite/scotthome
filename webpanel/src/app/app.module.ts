import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import {NotificationDisplayComponent} from './common/notification-display/notification-display.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatIconModule} from '@angular/material/icon';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatButtonToggleModule, MatListModule, MatSidenavModule, MatSliderModule} from '@angular/material';
import {MaterialModule} from './modules/material.module';
import {AppRoutingModule} from './app-routing.module';
import {LightRoomCardComponent} from './cards/lightroomcard/lightroomcard.component';
import {HttpClientModule} from '@angular/common/http';
import {CardComponent} from './cards/card/card.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    NotificationDisplayComponent,
    LightRoomCardComponent,
    CardComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatIconModule,
    MatToolbarModule,
    AppRoutingModule,
    HttpClientModule,
    MaterialModule,
    MatSidenavModule,
    MatListModule,
    MatSliderModule,
    MatButtonToggleModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {


}
