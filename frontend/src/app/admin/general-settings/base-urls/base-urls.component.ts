import {Component, Inject, OnInit} from '@angular/core';
import {GeneralBaseURLsUpdateRequest} from '../../types/types';
import {ToastNotificationService} from '../../../services/notification/toast-notification.service';
import {DefaultResponse, ErrorResponse, Fetcher} from '../../../interfaces/fetcher';
import {SessionStorageService} from '../../../services/session/session-storage.service';
import {ErrorHandlerService} from '../../../services/errors/error-handler.service';
import {ResponseStatus} from '../../../services/fetcher/statuses';
import {CodesService} from '../../services/errors/codes.service';
import {StorageService} from '../../services/storage/storage.service';

@Component({
  selector: 'app-base-urls',
  templateUrl: './base-urls.component.html',
  styleUrls: ['./base-urls.component.scss']
})
export class BaseUrlsComponent implements OnInit {
  public request: GeneralBaseURLsUpdateRequest = {
    base_url: '',
    command_hashes: [],
  };

  constructor(
    private readonly storage: StorageService,
    private readonly codesService: CodesService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly sessionStorage: SessionStorageService,
    private readonly toastNotification: ToastNotificationService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  public get updatingDisabled(): boolean {
    return this.request.command_hashes.length === 0;
  }

  public updateCommandsBaseURLs(): void {
    if (this.updatingDisabled) {
      this.toastNotification.info('You need to choose at least one command to update');
      return;
    }

    this.fetcher.patch('entity/mass-base-urls/', {
      account_hash: this.sessionStorage.getSessionId(),
      base_url: this.request.base_url,
      command_hashes: this.request.command_hashes.map(h => ({hash: h})),
    })
      .then(r => this.processMassUpdateResponse(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private processMassUpdateResponse(response: DefaultResponse | ErrorResponse): void {
    if (response.status === ResponseStatus.OK) {
      this.storage.invalidateCommands();
      this.request = {
        base_url: '',
        command_hashes: [],
      };
      this.toastNotification.success('Base URLs are updated successfully');
    } else {
      this.toastNotification.error(this.codesService.getMessageByDescription(
        (response as ErrorResponse).data.description
      ));
    }
  }

  ngOnInit(): void {
  }
}
