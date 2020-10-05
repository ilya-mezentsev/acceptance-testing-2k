import {Inject, Injectable} from '@angular/core';
import {CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, Router} from '@angular/router';
import {SessionStorageService} from './services/session/session-storage.service';
import {ErrorHandlerService} from './services/errors/error-handler.service';
import {Fetcher} from './interfaces/fetcher';
import {ResponseStatus} from './services/fetcher/statuses';

@Injectable({
  providedIn: 'root'
})
export class SessionGuard implements CanActivate {
  constructor(
    private readonly router: Router,
    private readonly errorHandler: ErrorHandlerService,
    private readonly sessionStorage: SessionStorageService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  async canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): Promise<boolean> {
    if (!!this.sessionStorage.hasSession()) {
      return Promise.resolve(true);
    } else {
      return this.hasSession();
    }
  }

  private async hasSession(): Promise<boolean> {
    let hasSession: boolean;
    try {
      const sessionResponse = await this.fetcher.get('session');
      hasSession = sessionResponse.status === ResponseStatus.OK;
      if (hasSession) {
        this.sessionStorage.saveSession(sessionResponse.data);
      }
    } catch (err) {
      this.errorHandler.handle(err);

      hasSession = false;
    }

    if (!hasSession) {
      this.router.navigate(['/public'])
        .catch(e => this.errorHandler.handle(e));
    }

    return hasSession;
  }
}
