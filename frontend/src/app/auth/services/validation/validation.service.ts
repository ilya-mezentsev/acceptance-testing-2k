import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ValidationService {
  private readonly minPasswordLength = 5;
  private readonly loginPattern = /^[-a-zA-Z0-9_]{1,64}$/;

  public validLogin(login: string): boolean {
    return this.loginPattern.test(login);
  }

  public validPassword(password: string): boolean {
    return password.length > this.minPasswordLength;
  }
}
