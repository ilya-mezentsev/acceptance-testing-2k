import {Component, Inject, OnInit} from '@angular/core';
import {MaterializeInitService} from '../../../services/materialize/materialize-init.service';
import {DefaultResponse, ErrorResponse, Fetcher, Response, ServerResponse} from '../../../interfaces/fetcher';
import {SessionStorageService} from '../../../services/session/session-storage.service';
import {ActivatedRoute, Router} from '@angular/router';
import {ErrorHandlerService} from '../../../services/errors/error-handler.service';
import {ResponseStatus} from '../../../services/fetcher/statuses';
import {ToastNotificationService} from '../../../services/notification/toast-notification.service';
import {CreateTestCommandResponse, KeyValueMapping, TestCommandMeta, TestCommandSettings} from '../../types/types';
import {StorageService} from '../../services/storage/storage.service';
import {CodesService} from '../../services/errors/codes.service';
import {HashService} from '../../../services/hash/hash.service';

@Component({
  selector: 'app-create-command',
  templateUrl: './create-command.component.html',
  styleUrls: ['./create-command.component.scss']
})
export class CreateCommandComponent implements OnInit {
  public headers: KeyValueMapping[] = [];
  public cookies: KeyValueMapping[] = [];
  public commandSettings: TestCommandSettings = {
    name: '',
    method: 'GET',
    pass_arguments_in_url: false,
    base_url: '',
    endpoint: '',
    timeout: 0,
    hash: '',
    object_hash: ''
  };
  private objectHash = '';
  private objectName = '';

  constructor(
    private readonly route: ActivatedRoute,
    private readonly router: Router,
    private readonly session: SessionStorageService,
    private readonly storage: StorageService,
    private readonly hashService: HashService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly codes: CodesService,
    private readonly materializeInit: MaterializeInitService,
    private readonly toastNotification: ToastNotificationService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) {}

  public hasHeaders(): boolean {
    return this.headers.length > 0;
  }

  public addHeader(): void {
    this.headers.push({key: '', value: '', hash: this.hashService.getRandomHash()} as any);
  }

  public removeHeader(hash: string): void {
    this.headers = this.headers.filter(h => h.hash !== hash);
  }

  public hasCookies(): boolean {
    return this.cookies.length > 0;
  }

  public addCookie(): void {
    this.cookies.push({key: '', value: '', hash: this.hashService.getRandomHash()} as any);
  }

  public removeCookie(hash: string): void {
    this.cookies = this.cookies.filter(c => c.hash !== hash);
  }

  public createCommand(): void {
    this.fetcher.post('entity/test-command/', {
      account_hash: this.session.getSessionId(),
      command_settings: {
        name: this.commandSettings.name,
        object_hash: this.objectHash,
        method: this.commandSettings.method,
        base_url: this.commandSettings.base_url,
        endpoint: this.commandSettings.endpoint,
        timeout: this.commandSettings.timeout,
        pass_arguments_in_url: this.commandSettings.pass_arguments_in_url
      }
    })
      .then(r => this.tryCreateCommandMeta(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private tryCreateCommandMeta(response: ServerResponse): void {
    if (response.status === ResponseStatus.OK) {
      if (this.hasHeaders() || this.hasCookies()) {
        const commandMeta: TestCommandMeta = {} as any;
        this.hasHeaders() && (commandMeta.headers = this.headers);
        this.hasCookies() && (commandMeta.cookies = this.cookies);

        this.fetcher.post('entity/test-command-meta/', {
          account_hash: this.session.getSessionId(),
          command_hash: (response as Response<CreateTestCommandResponse>).data.command_hash,
          command_meta: commandMeta
        })
          .then(r => this.processCreateCommandResponse(r))
          .catch(err => this.errorHandler.handle(err));
      } else {
       this.onSuccessCreation();
      }
    } else {
      this.toastNotification.error(this.codes.getMessageByDescription(
        (response as ErrorResponse).data.description
      ));
    }
  }

  private processCreateCommandResponse(response: DefaultResponse | ErrorResponse): void {
    if (response.status === ResponseStatus.OK) {
      this.onSuccessCreation();
    } else {
      this.toastNotification.error(this.codes.getMessageByDescription(
        (response as ErrorResponse).data.description
      ));
    }
  }

  private onSuccessCreation(): void {
    this.toastNotification.success('Command created successfully');
    this.storage.invalidateCommands();
    this.router.navigate(['/admin'])
      .catch(err => this.errorHandler.handle(err));
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.objectHash = params.get('object_hash');
      this.setCurrentObjectName();
    });

    this.materializeInit.initModals();
  }

  private setCurrentObjectName(): void {
    for (const object of this.storage.objects) {
      if (object.hash === this.objectHash) {
        this.objectName = object.name;
        return;
      }
    }

    this.router.navigate(['/admin'])
      .catch(err => this.errorHandler.handle(err));
  }
}
