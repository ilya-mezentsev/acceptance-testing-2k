import { Injectable } from '@angular/core';
import {Codes} from '../../../services/errors/codes';

@Injectable({
  providedIn: 'root'
})
export class CodesService extends Codes {
  constructor() {
    super();

    this.descriptionsToErrorMessage.set('unique-entity-exists', 'Object already exists');
  }
}
