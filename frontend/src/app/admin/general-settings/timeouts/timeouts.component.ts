import {Component, Inject, OnInit} from '@angular/core';
import {GeneralTimeoutUpdateRequest} from '../../types/types';
import {ToastNotificationService} from '../../../services/notification/toast-notification.service';
import {DefaultResponse, ErrorResponse, Fetcher} from '../../../interfaces/fetcher';
import {StorageService} from '../../services/storage/storage.service';
import {CodesService} from '../../services/errors/codes.service';
import {ErrorHandlerService} from '../../../services/errors/error-handler.service';
import {SessionStorageService} from '../../../services/session/session-storage.service';
import {ResponseStatus} from '../../../services/fetcher/statuses';

@Component({
  selector: 'app-timeouts',
  templateUrl: './timeouts.component.html',
  styleUrls: ['./timeouts.component.scss']
})
export class TimeoutsComponent implements OnInit {
  public request: GeneralTimeoutUpdateRequest = {
    timeout: 0,
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
    return this.request.command_hashes.length < 1;
  }

  public updateCommandsTimeout(): void {
    if (this.updatingDisabled) {
      this.toastNotification.info('You need to choose at least one command to update');
      return;
    }

    this.fetcher.patch('entity/mass-timeouts/', {
      account_hash: this.sessionStorage.getSessionId(),
      timeout: this.request.timeout,
      command_hashes: this.request.command_hashes.map(h => ({hash: h})),
    })
      .then(r => this.processMassUpdateResponse(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private processMassUpdateResponse(response: DefaultResponse | ErrorResponse): void {
    if (response.status === ResponseStatus.OK) {
      this.storage.invalidateCommands();
      this.request = {
        timeout: 0,
        command_hashes: [],
      };
      this.toastNotification.success('Timeouts are updated successfully');
    } else {
      this.toastNotification.error(this.codesService.getMessageByDescription(
        (response as ErrorResponse).data.description
      ));
    }
  }

  ngOnInit(): void {
  }
}
