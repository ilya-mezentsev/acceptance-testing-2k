import { Injectable } from '@angular/core';
import {ValidationRulesService} from '../../../services/validation/validation-rules.service';

@Injectable({
  providedIn: 'root'
})
export class ValidationService {
  constructor(
    private readonly validationRules: ValidationRulesService
  ) { }

  public isObjectNameValid(name: string): boolean {
    return this.validationRules.isRegularNameValid(name);
  }
}
