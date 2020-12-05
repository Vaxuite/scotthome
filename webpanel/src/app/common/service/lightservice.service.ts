import { Injectable } from '@angular/core';
import {HttpClient, HttpErrorResponse} from '@angular/common/http';
import {Observable, throwError} from 'rxjs';
import { environment } from 'src/environments/environment';

import {GroupModification, LightRoom, Scene} from '../domain/lights/room';
import {catchError} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class LightService {

  private readonly lightServiceBase = environment.lightServiceBase;
  private readonly serviceVersion = 'v1';
  private readonly path = `${this.lightServiceBase}/${this.serviceVersion}`;

  constructor(private http: HttpClient) { }

  public getRooms(): Observable<LightRoom[]> {
    return this.http.get<LightRoom[]>(`${this.path}/rooms`);
  }

  public setState(groupModification: GroupModification): Observable<any> {
    return this.http.post<any>(`${this.path}/rooms/state`, groupModification).pipe(
      catchError(this.handleError)
    );
  }

  public setBrightness(groupModification: GroupModification): Observable<any> {
    return this.http.post<any>(`${this.path}/rooms/bri`, groupModification).pipe(
      catchError(this.handleError)
    );
  }

  public setScene(groupModification: GroupModification): Observable<any> {
    return this.http.post<any>(`${this.path}/rooms/scene`, groupModification).pipe(
      catchError(this.handleError)
    );
  }

  private handleError(error: HttpErrorResponse) {
    if (error.error instanceof ErrorEvent) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error.message);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong.
      console.error(
        `Backend returned code ${error.status}, ` +
        `body was: ${error.error}`);
    }
    // Return an observable with a user-facing error message.
    return throwError(
      'Something bad happened; please try again later.');
  }
}
