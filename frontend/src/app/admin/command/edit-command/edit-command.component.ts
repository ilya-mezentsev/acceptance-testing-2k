import {Component, Inject, OnInit} from '@angular/core';
import {ErrorHandlerService} from '../../../services/errors/error-handler.service';
import {SessionStorageService} from '../../../services/session/session-storage.service';
import {MaterializeInitService} from '../../../services/materialize/materialize-init.service';
import {ToastNotificationService} from '../../../services/notification/toast-notification.service';
import {DefaultResponse, ErrorResponse, Fetcher, Response, UpdatePayload} from '../../../interfaces/fetcher';
import {ActivatedRoute, Router} from '@angular/router';
import {StorageService} from '../../services/storage/storage.service';
import {CodesService} from '../../services/errors/codes.service';
import {
  KeyValueMapping,
  TestCommandMeta,
  TestCommandRecord,
  TestCommandSettings
} from '../../types/types';
import {HashService} from '../../../services/hash/hash.service';
import {ResponseStatus} from '../../../services/fetcher/statuses';

@Component({
  selector: 'app-edit-command',
  templateUrl: './edit-command.component.html',
  styleUrls: ['./edit-command.component.scss']
})
export class EditCommandComponent implements OnInit {

  constructor(
    private readonly router: Router,
    private readonly route: ActivatedRoute,
    private readonly storage: StorageService,
    private readonly hashService: HashService,
    private readonly codesService: CodesService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly sessionStorage: SessionStorageService,
    private readonly materializeInit: MaterializeInitService,
    private readonly toastNotification: ToastNotificationService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  public get headers(): KeyValueMapping[] {
    return this.newHeaders.concat(this.existsHeaders);
  }

  public get cookies(): KeyValueMapping[] {
    return this.newCookies.concat(this.existsCookies);
  }
  public commandSettings: TestCommandSettings = {} as any;
  private commandHash = '';
  private newHeaders: KeyValueMapping[] = [];
  private newCookies: KeyValueMapping[] = [];
  private existsHeaders: KeyValueMapping[] = [];
  private existsCookies: KeyValueMapping[] = [];
  private currentCommand: TestCommandRecord = {} as any;
  private commandSettingsWereChanged = false;
  private commandMetaWasChanged = false;
  private commandMetaWereAdded = false;
  private readonly editableCommandSettingFields = [
    'method', 'base_url', 'endpoint', 'timeout', 'pass_arguments_in_url'
  ];

  private static getDiff(
    storedKeyValues: KeyValueMapping[],
    keyValues: KeyValueMapping[],
  ): UpdatePayload[] {
    return keyValues.reduce((updatePayload, keyValue) => {
      const storedKeyValue = storedKeyValues.find(h => h.hash === keyValue.hash);

      if (keyValue.key !== storedKeyValue.key) {
        updatePayload.push({
          hash: keyValue.hash,
          field_name: 'key',
          new_value: keyValue.key,
        });
      }

      if (keyValue.value !== storedKeyValue.value) {
        updatePayload.push({
          hash: keyValue.hash,
          field_name: 'value',
          new_value: keyValue.value,
        });
      }

      return updatePayload;
    }, [] as UpdatePayload[]);
  }

  public addHeader(): void {
    this.newHeaders.push(
      {key: '', value: '', hash: this.hashService.getRandomHash()} as any
    );
  }

  public removeHeader(hash: string): void {
    if (this.newHeaders.find(h => h.hash === hash)) {
      this.newHeaders = this.newHeaders.filter(h => h.hash !== hash);
    } else {
      this.removeExistsHeader(hash);
    }
  }

  private removeExistsHeader(hash: string): void {
    this.fetcher
      .delete(`entity/test-command-headers/${this.sessionStorage.getSessionId()}/${hash}/`)
      .then(r => this.processRequest(r))
      .then(() => this.currentCommand.headers = this.currentCommand.headers.filter(h => h.hash !== hash))
      .then(() => this.initCommandSettingsAndMeta())
      .catch(err => this.errorHandler.handle(err));
  }

  private processRequest(response: DefaultResponse | ErrorResponse): Promise<void> {
    if (response.status === ResponseStatus.OK) {
      return Promise.resolve();
    } else {
      this.toastNotification.error(this.codesService.getMessageByDescription(
        (response as ErrorResponse).data.description
      ));
      return Promise.reject();
    }
  }

  public addCookie(): void {
    this.newCookies.push({key: '', value: '', hash: this.hashService.getRandomHash()} as any);
  }

  public removeCookie(hash: string): void {
    if (this.newCookies.find(c => c.hash === hash)) {
      this.newCookies = this.newCookies.filter(c => c.hash !== hash);
    } else {
      this.removeExistsCookie(hash);
    }
  }

  private removeExistsCookie(hash: string): void {
    this.fetcher
      .delete(`entity/test-command-cookies/${this.sessionStorage.getSessionId()}/${hash}/`)
      .then(r => this.processRequest(r))
      .then(() => this.currentCommand.cookies = this.currentCommand.cookies.filter(c => c.hash !== hash))
      .then(() => this.initCommandSettingsAndMeta())
      .catch(err => this.errorHandler.handle(err));
  }

  public saveCommand(): void {
    Promise.all([
      this.updateCommandSettings(),
      this.updateCommandMeta(),
      this.createNewCommandMeta(),
    ])
      .then(() => this.postProcessUpdating())
      .catch(err => this.errorHandler.handle(err));
  }

  private updateCommandSettings(): Promise<void> {
    this.commandSettingsWereChanged =
      this.commandSettings.name !== this.currentCommand.name ||
      this.editableCommandSettingFields.some(
        fieldName => this.commandSettings[fieldName] !== this.currentCommand[fieldName]
      );

    if (this.commandSettingsWereChanged) {
      return this.fetcher.patch(`entity/test-command/`, {
        account_hash: this.sessionStorage.getSessionId(),
        exists_command: {...this.currentCommand},
        updated_command: {...this.commandSettings}
      }).then(r => this.processRequest(r));
    } else {
      return Promise.resolve();
    }
  }

  private updateCommandMeta(): Promise<void> {
    const headersUpdatePayload: UpdatePayload[] =
      EditCommandComponent.getDiff(this.currentCommand.headers, this.existsHeaders);

    const cookiesUpdatePayload: UpdatePayload[] =
      EditCommandComponent.getDiff(this.currentCommand.cookies, this.existsCookies);

    const updatePayload = {};
    const headersChanged = headersUpdatePayload.length > 0;
    const cookiesChanged = cookiesUpdatePayload.length > 0;
    headersChanged && (updatePayload['headers'] = headersUpdatePayload);
    cookiesChanged && (updatePayload['cookies'] = cookiesUpdatePayload);

    if (headersChanged || cookiesChanged) {
      this.commandSettingsWereChanged = true;
      return this.fetcher.patch(`entity/test-command-meta/`, {
        account_hash: this.sessionStorage.getSessionId(),
        ...updatePayload
      }).then(r => this.processRequest(r));
    } else {
      return Promise.resolve();
    }
  }

  private createNewCommandMeta(): Promise<void> {
    const hasNewHeaders = this.newHeaders.length > 0;
    const hasNewCookies = this.newCookies.length > 0;
    this.commandMetaWereAdded = hasNewHeaders || hasNewCookies;
    if (!this.commandMetaWereAdded) {
      return Promise.resolve();
    }

    const commandMeta: TestCommandMeta = {} as any;
    hasNewHeaders && (commandMeta.headers = this.newHeaders);
    hasNewCookies && (commandMeta.cookies = this.newCookies);

    return this.fetcher.post('entity/test-command-meta/', {
      account_hash: this.sessionStorage.getSessionId(),
      command_hash: this.currentCommand.hash,
      command_meta: commandMeta
    }).then(r => this.processRequest(r));
  }

  private postProcessUpdating(): void {
    if (
      this.commandSettingsWereChanged ||
      this.commandMetaWasChanged ||
      this.commandMetaWereAdded
    ) {
      this.toastNotification.info('Command updated successfully');
      this.storage.invalidateCommands();
    } else {
      this.toastNotification.info('Command did not changed so not updated');
    }

    this.router.navigate(['/admin'])
      .catch(err => this.errorHandler.handle(err));
  }

  public deleteCommand(): void {
    this.fetcher.delete(`entity/test-command/${
      this.sessionStorage.getSessionId()}/${this.currentCommand.hash}/`)
      .then(r => this.processRequest(r))
      .then(() => {
        this.storage.invalidateCommands();
        this.toastNotification.success('Command deleted successfully');
        return this.router.navigate(['/admin']);
      })
      .catch(err => this.errorHandler.handle(err));
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.commandHash = params.get('command_hash');

      this.initCurrentCommand();
      this.initCommandSettingsAndMeta();
    });

    this.materializeInit.initModals();
  }

  private initCurrentCommand(): void {
    const currentCommand = this.storage.commands.find(c => c.hash === this.commandHash);
    if (currentCommand) {
      this.currentCommand = currentCommand;
    } else {
      this.tryFetchCurrentCommand();
    }
  }

  private tryFetchCurrentCommand(): void {
    this.fetcher
      .get(`entity/test-command/${this.sessionStorage.getSessionId()}/${this.commandHash}/`)
      .then(r => {
        if (r.status === ResponseStatus.OK) {
          this.currentCommand = (r as Response<TestCommandRecord>).data;
          this.initCommandSettingsAndMeta();
        } else {
          this.toastNotification.error('Unable to fetch test command');
          return Promise.reject();
        }
      })
      .catch(err => {
        this.errorHandler.handle(err);
        this.router.navigate(['/admin'])
          .catch(e => this.errorHandler.handle(e));
      });
  }

  private initCommandSettingsAndMeta(): void {
    this.commandSettings = {
      name: this.currentCommand.name,
      hash: this.currentCommand.hash,
      object_hash: this.currentCommand.object_hash,
      method: this.currentCommand.method,
      base_url: this.currentCommand.base_url,
      endpoint: this.currentCommand.endpoint,
      timeout: this.currentCommand.timeout,
      pass_arguments_in_url: this.currentCommand.pass_arguments_in_url,
    };

    // we need to copy headers and cookies to detect changes and perform update later
    this.existsHeaders = this.currentCommand.headers?.map(h => ({...h})) || [];
    this.existsCookies = this.currentCommand.cookies?.map(c => ({...c})) || [];
  }
}
