import {Component, Inject, OnInit} from '@angular/core';
import {ValidationService} from '../services/validation/validation.service';
import {DefaultResponse, ErrorResponse, Fetcher} from '../../interfaces/fetcher';
import {ErrorHandlerService} from '../../services/errors/error-handler.service';
import {ToastNotificationService} from '../../services/notification/toast-notification.service';
import {ResponseStatus} from '../../services/fetcher/statuses';
import {NavigationService} from '../services/navigation/navigation.service';
import {CodesService} from '../services/errors/codes.service';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.scss']
})
export class RegistrationComponent implements OnInit {
  private login = '';
  private password = '';
  private repeatedPassword = '';

  constructor(
    private readonly validation: ValidationService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly toastNotification: ToastNotificationService,
    private readonly authNavigator: NavigationService,
    private readonly errorDescriptionDecoder: CodesService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) {}

  public changeLogin(value: string): void {
    this.login = value;
  }

  public changePassword(value: string): void {
    this.password = value;
  }

  public changeRepeatedPassword(value: string): void {
    this.repeatedPassword = value;
  }

  public get registrationDisabled(): boolean {
    return !this.validation.validLogin(this.login) ||
      !this.validation.validPassword(this.password);
  }

  private get passwordsArentMatch(): boolean {
    return this.password != this.repeatedPassword;
  }

  public tryRegister(): void {
    if (this.passwordsArentMatch) {
      this.toastNotification.info('Passwords are not match');
      return;
    }

    if (this.registrationDisabled) {
      this.toastNotification.info('You need to enter valid data to register');
      return;
    }

    this.fetcher.post('/registration/', {
      login: this.login,
      password: this.password
    })
      .then(r => this.processResponse(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private processResponse(response: DefaultResponse | ErrorResponse): void {
    if (response.status === ResponseStatus.OK) {
      this.toastNotification.success('Registration successfully performed!');
      this.authNavigator.navigateToLogin();
    } else {
      this.toastNotification.error(
        this.errorDescriptionDecoder.getMessageByDescription(response.data.description)
      );
    }
  }

  ngOnInit(): void {
  }
}
