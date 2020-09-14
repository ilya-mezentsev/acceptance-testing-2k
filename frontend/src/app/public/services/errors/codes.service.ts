import { Injectable } from '@angular/core';
import {Codes} from '../../../services/errors/codes';

@Injectable({
  providedIn: 'root'
})
export class CodesService extends Codes {
  constructor() {
    super();

    [
      ['login-already-exists', 'Entered login already exists'],
      ['account-does-not-exists', 'Account does not exists']
    ].forEach(descriptionToErrorMessage => {
      this.descriptionsToErrorMessage.set(
        descriptionToErrorMessage[0],
        descriptionToErrorMessage[1]
      );
    });
  }
}
