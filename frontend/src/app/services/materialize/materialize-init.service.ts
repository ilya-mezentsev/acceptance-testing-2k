import { Injectable } from '@angular/core';

declare namespace M {
  const Modal: any;
}

@Injectable({
  providedIn: 'root'
})
export class MaterializeInitService {
  public initModals(): void {
    M.Modal.init(document.querySelectorAll('.modal'), {});
  }
}
