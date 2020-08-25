import { Injectable } from '@angular/core';
import {CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, Router} from '@angular/router';
import {SessionStorageService} from './services/session/session-storage.service';
import {ErrorHandlerService} from './services/errors/error-handler.service';

@Injectable({
  providedIn: 'root'
})
export class SessionGuard implements CanActivate {
  constructor(
    private readonly router: Router,
    private readonly errorHandler: ErrorHandlerService,
    private readonly sessionStorage: SessionStorageService,
  ) { }

  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
    if (!!this.sessionStorage.hasSession()) {
      return true;
    } else {
      this.router.navigate(['/authorization'])
        .catch(err => this.errorHandler.handle(err));
      return false;
    }
  }
}
