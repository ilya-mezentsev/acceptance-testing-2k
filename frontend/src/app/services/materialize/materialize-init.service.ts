import { Injectable } from '@angular/core';

declare namespace M {
  const Modal: any;
  const FormSelect: any;
  const Collapsible: any;
}

@Injectable({
  providedIn: 'root'
})
export class MaterializeInitService {
  public initModals(): void {
    M.Modal.init(document.querySelectorAll('.modal'), {});
  }

  public initSelects(): void {
    M.FormSelect.init(document.querySelectorAll('select'), {});
  }

  public initCollapsible(): void {
    M.Collapsible.init(document.querySelectorAll('.collapsible.popout'), {
      accordion: false
    });
  }
}
