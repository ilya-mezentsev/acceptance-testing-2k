import {Component, Inject, OnInit} from '@angular/core';
import {MaterializeInitService} from '../../services/materialize/materialize-init.service';
import {CreateTestCommandResponse, KeyValueMapping, TestCommandMeta} from '../types/types';
import {DefaultResponse, ErrorResponse, Fetcher, Response, ServerResponse} from '../../interfaces/fetcher';
import {SessionStorageService} from '../../services/session/session-storage.service';
import {ActivatedRoute, Router} from '@angular/router';
import {ErrorHandlerService} from '../../services/errors/error-handler.service';
import {StorageService} from '../services/storage/storage.service';
import {ResponseStatus} from '../../services/fetcher/statuses';
import {ToastNotificationService} from '../../services/notification/toast-notification.service';
import {CodesService} from '../services/errors/codes.service';

@Component({
  selector: 'app-create-command',
  templateUrl: './create-command.component.html',
  styleUrls: ['./create-command.component.scss']
})
export class CreateCommandComponent implements OnInit {
  public passArgumentsInURL = false;
  public method = 'GET';
  public headers: KeyValueMapping[] = [];
  public cookies: KeyValueMapping[] = [];
  private commandName: string;
  private baseURL: string;
  private endpoint: string;
  private objectHash = '';
  private objectName = '';

  constructor(
    private readonly route: ActivatedRoute,
    private readonly router: Router,
    private readonly session: SessionStorageService,
    private readonly storage: StorageService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly codes: CodesService,
    private readonly materializeInit: MaterializeInitService,
    private readonly toastNotification: ToastNotificationService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) {}

  public setCommandName(value: string): void {
    this.commandName = value;
  }

  public setBaseURL(value: string): void {
    this.baseURL = value;
  }

  public setEndpoint(value: string): void {
    this.endpoint = value;
  }

  public hasHeaders(): boolean {
    return this.headers.length > 0;
  }

  public addHeader(): void {
    this.headers.push({key: '', value: ''});
  }

  public removeHeader(index: number): void {
    this.headers = this.headers.filter((_, i) => i !== index);
  }

  public hasCookies(): boolean {
    return this.cookies.length > 0;
  }

  public addCookie(): void {
    this.cookies.push({key: '', value: ''});
  }

  public removeCookie(index: number): void {
    this.cookies = this.cookies.filter((_, i) => i !== index);
  }

  public createCommand(): void {
    this.fetcher.post('entity/test-command/', {
      account_hash: this.session.getSessionId(),
      command_settings: {
        name: this.commandName,
        object_name: this.objectName,
        method: this.method,
        base_url: this.baseURL,
        endpoint: this.endpoint,
        pass_arguments_in_url: this.passArgumentsInURL
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
    this.materializeInit.initSelects();

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
