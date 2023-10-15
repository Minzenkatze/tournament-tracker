import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class CompetitorService {
  constructor(private http: HttpClient) {}

  getCompetitors(): any {
    console.log('hello');
    return this.http.get<any>('localhost:8080/competitors');
  }
}

export interface Competitor {
  id: number;
  firstName: string;
  lastName: string;
  weight: string;
}
