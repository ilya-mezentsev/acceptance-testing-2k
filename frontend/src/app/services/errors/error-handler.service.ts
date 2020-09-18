import { Injectable } from '@angular/core';
import {MaterializeInitService} from "../materialize/materialize-init.service";

@Injectable({
  providedIn: 'root'
})
export class ErrorHandlerService {
  private readonly unexpectedErrorModalId = 'unexpected-error';

  constructor(
    private readonly materializeInit: MaterializeInitService,
  ) { }

  public handle(err: any): void {
    this.materializeInit.initModals();
    this.materializeInit.showModalWithId(this.unexpectedErrorModalId);
    console.log(err);
  }
}
