import {Component, Inject, OnInit} from '@angular/core';
import {ValidationService} from '../services/validation/validation.service';
import {ErrorResponse, Fetcher, ServerResponse} from '../../interfaces/fetcher';
import {ErrorHandlerService} from '../../services/errors/error-handler.service';
import {SessionStorageService} from '../../services/session/session-storage.service';
import {ToastNotificationService} from '../../services/notification/toast-notification.service';
import {NavigationService} from '../services/navigation/navigation.service';
import {CodesService} from '../services/errors/codes.service';
import {ResponseStatus} from '../../services/fetcher/statuses';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  public login = '';
  public password = '';

  constructor(
    private readonly validation: ValidationService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly sessionStorage: SessionStorageService,
    private readonly toastNotification: ToastNotificationService,
    private readonly authNavigator: NavigationService,
    private readonly errorDescriptionDecoder: CodesService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  public changeLogin(value: string): void {
    this.login = value;
  }

  public changePassword(value: string): void {
    this.password = value;
  }

  public get loginDisabled(): boolean {
    return !this.validation.validLogin(this.login) ||
      !this.validation.validPassword(this.password);
  }

  public tryLogin(): void {
    if (this.loginDisabled) {
      this.toastNotification.info('You need to enter valid data to login');
      return;
    }

    this.fetcher.post('session/', {
      login: this.login,
      password: this.password
    })
      .then(r => this.processResponse(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private processResponse(response: ServerResponse): void {
    if (response.status === ResponseStatus.OK) {
      this.sessionStorage.saveSession(response.data);
      this.authNavigator.navigateToAuthorized();
    } else {
      this.toastNotification.error(
        this.errorDescriptionDecoder.getMessageByDescription(
          (response as ErrorResponse).data.description
        )
      );
    }
  }

  ngOnInit(): void {
  }
}
