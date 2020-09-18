import { Injectable } from '@angular/core';
import {Codes} from '../../../services/errors/codes';

@Injectable({
  providedIn: 'root'
})
export class CodesService extends Codes {
  constructor() {
    super();

    [
      ['unique-entity-exists', 'Object already exists'],
      ['no-test-cases', 'No tests in sent file'],
    ].forEach(descriptionToErrorMessage => {
      this.descriptionsToErrorMessage.set(
        descriptionToErrorMessage[0],
        descriptionToErrorMessage[1]
      );
    });
  }
}
