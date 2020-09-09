import {Component, Inject, OnInit} from '@angular/core';
import {GeneralHeadersCreateRequest} from '../../types/types';
import {HashService} from '../../../services/hash/hash.service';
import {ToastNotificationService} from '../../../services/notification/toast-notification.service';
import {StorageService} from '../../services/storage/storage.service';
import {CodesService} from '../../services/errors/codes.service';
import {ErrorHandlerService} from '../../../services/errors/error-handler.service';
import {SessionStorageService} from '../../../services/session/session-storage.service';
import {DefaultResponse, ErrorResponse, Fetcher} from '../../../interfaces/fetcher';
import {ResponseStatus} from '../../../services/fetcher/statuses';

@Component({
  selector: 'app-headers',
  templateUrl: './headers.component.html',
  styleUrls: ['./headers.component.scss']
})
export class HeadersComponent implements OnInit {
  public request: GeneralHeadersCreateRequest = {
    headers: [],
    command_hashes: []
  };

  constructor(
    private readonly storage: StorageService,
    private readonly hashService: HashService,
    private readonly codesService: CodesService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly sessionStorage: SessionStorageService,
    private readonly toastNotification: ToastNotificationService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  public get addingDisabled(): boolean {
    return (
      this.request.headers.length < 1 ||
      this.request.command_hashes.length < 1
    );
  }

  public addEmptyHeader(): void {
    this.request.headers.push({
      key: '',
      value: '',
      hash: this.hashService.getRandomHash()
    } as any);
  }

  public removeHeader(hash: string): void {
    this.request.headers = this.request.headers.filter(h => h.hash !== hash);
  }

  public createHeaders(): void {
    if (this.addingDisabled) {
      this.toastNotification.info('You need to add header of choose command');
      return;
    }

    this.fetcher.post('entity/mass-headers/', {
      account_hash: this.sessionStorage.getSessionId(),
      headers: this.request.headers,
      command_hashes: this.request.command_hashes.map(h => ({hash: h}))
    })
      .then(r => this.processMassCreateResponse(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private processMassCreateResponse(response: DefaultResponse | ErrorResponse): void {
    if (response.status === ResponseStatus.OK) {
      this.storage.invalidateCommands();
      this.request = {
        headers: [],
        command_hashes: [],
      };
      this.toastNotification.success('Headers are successfully created');
    } else {
      this.toastNotification.error(this.codesService.getMessageByDescription(
        (response as ErrorResponse).data.description
      ));
    }
  }

  ngOnInit(): void {
  }
}
