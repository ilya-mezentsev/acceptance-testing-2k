import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ErrorHandlerService {
  public handle(err: any): void {
    console.log(err);
  }
}
