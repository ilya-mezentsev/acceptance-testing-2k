import {Component, Inject, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {StorageService} from '../services/storage/storage.service';
import {ErrorHandlerService} from '../../services/errors/error-handler.service';
import {KeyValueMapping, TestCommandRecord, TestObject} from '../types/types';
import {ValidationService} from '../services/validation/validation.service';
import {ToastNotificationService} from '../../services/notification/toast-notification.service';
import {DefaultResponse, ErrorResponse, Fetcher, Response} from '../../interfaces/fetcher';
import {ResponseStatus} from '../../services/fetcher/statuses';
import {CodesService} from '../services/errors/codes.service';
import {MaterializeInitService} from '../../services/materialize/materialize-init.service';

@Component({
  selector: 'app-edit-object',
  templateUrl: './edit-object.component.html',
  styleUrls: ['./edit-object.component.scss']
})
export class EditObjectComponent implements OnInit {
  public objectHash = '';
  public objectName = '';
  public commands: TestCommandRecord[] = [];
  private currentObject: TestObject;

  constructor(
    private readonly router: Router,
    private readonly route: ActivatedRoute,
    private readonly storage: StorageService,
    private readonly codesService: CodesService,
    private readonly validation: ValidationService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly materializeInit: MaterializeInitService,
    private readonly toastNotification: ToastNotificationService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  public hasCommands(): boolean {
    return this.commands.length > 0;
  }

  public setCurrentObjectName(name: string): void {
    this.objectName = name;
  }

  public reduceKeyValue(mapping: KeyValueMapping[]): string {
    if (!mapping?.length || mapping.length < 1) {
      return 'No values provided';
    }

    return mapping.map(m => `${m.key}=${m.value}`).join(';');
  }

  public get updatingDisabled(): boolean {
    return !this.validation.isObjectNameValid(this.objectName);
  }

  public updateObject(): void {
    if (this.updatingDisabled) {
      this.toastNotification.info('You need to enter valid object name');
      return;
    }

    if (this.objectName === this.currentObject.name) {
      this.toastNotification.info('Object name did not changed so not updated');
      this.router.navigate(['/admin'])
        .catch(err => this.errorHandler.handle(err));
      return;
    }

    this.fetcher.patch('test-object', {
      update_payload: [
        {
          hash: this.currentObject.hash,
          field_name: 'name',
          new_value: this.objectName
        }
      ]
    })
      .then(r => this.processUpdateResponse(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private processUpdateResponse(response: DefaultResponse | ErrorResponse): void {
    this.processResponse(
      response,
      'Object updated successfully'
    );
  }

  private processResponse(
    response: DefaultResponse | ErrorResponse,
    successMessage: string
  ): void {
    if (response.status === ResponseStatus.OK) {
      this.storage.invalidateObjects();
      this.toastNotification.success(successMessage);
      this.router.navigate(['/admin'])
        .catch(err => this.errorHandler.handle(err));
    } else {
      this.toastNotification.error(this.codesService.getMessageByDescription(
        (response as ErrorResponse).data.description
      ));
    }
  }

  public deleteObject(): void {
    this.fetcher.delete(`test-object/${this.currentObject.hash}`)
      .then(r => this.processDeleteResponse(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private processDeleteResponse(response: DefaultResponse | ErrorResponse): void {
    this.processResponse(
      response,
      'Object deleted successfully'
    );
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.objectHash = params.get('object_hash');
      this.initCurrentObject();
    });

    this.materializeInit.initModals();
    this.materializeInit.initCollapsible();
  }

  private initCurrentObject(): void {
    const currentObject = this.storage.objects.find(o => o.hash === this.objectHash);
    if (currentObject) {
      this.currentObject = currentObject;
      this.objectName = currentObject.name;
      this.setCurrentCommands();
    } else {
      this.tryFetchCurrentObjectAndCommands();
    }
  }

  private tryFetchCurrentObjectAndCommands(): void {
    this.fetcher
      .get(`test-object/${this.objectHash}`)
      .then(r => {
        if (r.status === ResponseStatus.OK) {
          this.currentObject = (r as Response<TestObject>).data;
          this.objectName = this.currentObject.name;
          return this.fetcher.get(`test-command`);
        } else {
          this.toastNotification.error('Unable to fetch test object');
          return Promise.reject();
        }
      })
      .then(r => {
        if (r.status === ResponseStatus.OK) {
          this.storage.commands = (r as Response<TestCommandRecord[]>).data;
          this.setCurrentCommands();
        } else {
          this.toastNotification.error('Unable to fetch test commands');
          return Promise.reject();
        }
      })
      .catch(err => {
        this.errorHandler.handle(err);
        this.router.navigate(['/admin'])
          .catch(e => this.errorHandler.handle(e));
      });
  }

  private setCurrentCommands(): void {
    this.commands = this.storage.commands.filter(c => c.object_hash === this.objectHash);
  }
}
