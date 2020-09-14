import { Injectable } from '@angular/core';
import {ValidationRulesService} from '../../../services/validation/validation-rules.service';

@Injectable({
  providedIn: 'root'
})
export class ValidationService {
  private readonly minPasswordLength = 5;

  constructor(
    private readonly validationRules: ValidationRulesService
  ) {}

  public validLogin(login: string): boolean {
    return this.validationRules.isRegularNameValid(login);
  }

  public validPassword(password: string): boolean {
    return password.length > this.minPasswordLength;
  }
}
