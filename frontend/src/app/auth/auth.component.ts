import {Component, Inject, OnInit} from '@angular/core';
import {Fetcher, ServerResponse} from '../interfaces/fetcher';
import {ErrorHandlerService} from '../services/errors/error-handler.service';
import {NavigationService} from './services/navigation/navigation.service';
import {SessionStorageService} from '../services/session/session-storage.service';
import {ResponseStatus} from '../services/fetcher/statuses';
import {CodesService} from './services/errors/codes.service';
import {ToastNotificationService} from '../services/notification/toast-notification.service';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss']
})
export class AuthComponent implements OnInit {
  constructor(
    private readonly errorHandler: ErrorHandlerService,
    private readonly authNavigator: NavigationService,
    private readonly sessionStorage: SessionStorageService,
    private readonly errorDescriptionDecoder: CodesService,
    private readonly toastNotification: ToastNotificationService,
    @Inject('Fetcher') private readonly fetcher: Fetcher
  ) {
    this.fetcher.get('session/')
      .then(r => this.navigateToAuthorizedIfNeeded(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private navigateToAuthorizedIfNeeded(response: ServerResponse): void {
    if (response.status === ResponseStatus.OK) {
      this.sessionStorage.saveSession(response.data);
      this.authNavigator.navigateToAuthorized();
    }
  }

  ngOnInit(): void {
  }
}
