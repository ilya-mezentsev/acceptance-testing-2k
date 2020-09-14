import { Injectable } from '@angular/core';

declare namespace M {
  const Modal: any;
  const FormSelect: any;
  const Collapsible: any;
  const Tooltip: any;
  const Sidenav: any;
  const Dropdown: any;
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

  public initCollapsibleWithoutOptions(): void {
    M.Collapsible.init(document.querySelectorAll('.collapsible'), {});
  }

  public initTooltips(): void {
    M.Tooltip.init(document.querySelectorAll('.with-tooltip'), {});
  }

  public initSidenav(): void {
    M.Sidenav.init(document.querySelectorAll('.sidenav'), {});
  }

  public initDropdowns(): void {
    M.Dropdown.init(document.querySelectorAll('.dropdown-trigger'), {});
  }
}
