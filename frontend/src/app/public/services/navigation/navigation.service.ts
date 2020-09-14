import { Injectable } from '@angular/core';
import {Router} from '@angular/router';
import {ErrorHandlerService} from '../../../services/errors/error-handler.service';

@Injectable({
  providedIn: 'root'
})
export class NavigationService {
  constructor(
    private readonly router: Router,
    private readonly errorHandler: ErrorHandlerService
  ) {}

  public navigateToLogin(): void {
    this.router.navigate(['/public/sign-in'])
      .catch(err => this.errorHandler.handle(err));
  }

  public navigateToAuthorized(): void {
    this.router.navigate(['/admin'])
      .catch(err => this.errorHandler.handle(err));
  }
}
