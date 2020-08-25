import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ValidationRulesService {
  private readonly regularNamePattern = /^[-a-zA-Z0-9_]{1,64}$/;

  public isRegularNameValid(name: string): boolean {
    return this.regularNamePattern.test(name);
  }
}
