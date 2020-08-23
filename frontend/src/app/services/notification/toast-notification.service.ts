import { Injectable } from '@angular/core';

declare namespace M {
  const toast: any;
}

@Injectable({
  providedIn: 'root'
})
export class ToastNotificationService {
  public success(message: string): void {
    M.toast({
      html: message,
      classes: 'toast toast-success'
    });
  }

  public info(message: string): void {
    M.toast({
      html: message,
      classes: 'toast toast-info'
    });
  }

  public error(message: string): void {
    M.toast({
      html: message,
      classes: 'toast toast-error'
    });
  }
}
