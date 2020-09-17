import {Component, Inject, OnInit} from '@angular/core';
import {GeneralCookiesCreateRequest} from '../../types/types';
import {HashService} from '../../../services/hash/hash.service';
import {ToastNotificationService} from '../../../services/notification/toast-notification.service';
import {StorageService} from '../../services/storage/storage.service';
import {CodesService} from '../../services/errors/codes.service';
import {ErrorHandlerService} from '../../../services/errors/error-handler.service';
import {DefaultResponse, ErrorResponse, Fetcher} from '../../../interfaces/fetcher';
import {ResponseStatus} from '../../../services/fetcher/statuses';

@Component({
  selector: 'app-cookies',
  templateUrl: './cookies.component.html',
  styleUrls: ['./cookies.component.scss']
})
export class CookiesComponent implements OnInit {
  public request: GeneralCookiesCreateRequest = {
    cookies: [],
    command_hashes: [],
  };

  constructor(
    private readonly storage: StorageService,
    private readonly hashService: HashService,
    private readonly codesService: CodesService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly toastNotification: ToastNotificationService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  public get addingDisabled(): boolean {
    return (
      this.request.command_hashes.length < 1 ||
      this.request.cookies.length < 1
    );
  }

  public addEmptyCookie(): void {
    this.request.cookies.push({
      key: '',
      value: '',
      hash: this.hashService.getRandomHash()
    } as any);
  }

  public removeCookie(hash: string): void {
    this.request.cookies = this.request.cookies.filter(c => c.hash !== hash);
  }

  public createCookies(): void {
    if (this.addingDisabled) {
      this.toastNotification.info('You need to add cookie or choose command');
      return;
    }

    this.fetcher.post('mass-cookies', {
      cookies: this.request.cookies,
      command_hashes: this.request.command_hashes.map(h => ({hash: h}))
    })
      .then(r => this.processMassCreateResponse(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private processMassCreateResponse(response: DefaultResponse | ErrorResponse): void {
    if (response.status === ResponseStatus.OK) {
      this.storage.invalidateCommands();
      this.request = {
        cookies: [],
        command_hashes: [],
      };
      this.toastNotification.success('Cookies are successfully created');
    } else {
      this.toastNotification.error(this.codesService.getMessageByDescription(
        (response as ErrorResponse).data.description
      ));
    }
  }

  ngOnInit(): void {
  }
}
