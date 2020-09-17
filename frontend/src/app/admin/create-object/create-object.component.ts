import {Component, Inject, OnInit} from '@angular/core';
import {ValidationService} from '../services/validation/validation.service';
import {ToastNotificationService} from '../../services/notification/toast-notification.service';
import {DefaultResponse, ErrorResponse, Fetcher} from '../../interfaces/fetcher';
import {Router} from '@angular/router';
import {ErrorHandlerService} from '../../services/errors/error-handler.service';
import {ResponseStatus} from '../../services/fetcher/statuses';
import {CodesService} from '../services/errors/codes.service';
import {StorageService} from '../services/storage/storage.service';

@Component({
  selector: 'app-create-object',
  templateUrl: './create-object.component.html',
  styleUrls: ['./create-object.component.scss']
})
export class CreateObjectComponent implements OnInit {
  private objectName = '';

  constructor(
    private readonly router: Router,
    private readonly validation: ValidationService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly toastNotification: ToastNotificationService,
    private readonly codesService: CodesService,
    private readonly storage: StorageService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  public changeObjectName(value: string): void {
    this.objectName = value;
  }

  public get creatingDisabled(): boolean {
    return !this.validation.isObjectNameValid(this.objectName);
  }

  public createObject(): void {
    if (this.creatingDisabled) {
      this.toastNotification.info('You need to enter valid object name');
      return;
    }

    this.fetcher.post('test-object', {
      test_object: {
        name: this.objectName
      }
    })
      .then(r => this.processResponse(r))
      .catch(err => this.errorHandler.handle(err));
  }

  private processResponse(response: DefaultResponse | ErrorResponse): void {
    if (response.status === ResponseStatus.OK) {
      this.storage.invalidateObjects();
      this.toastNotification.success('Object is successfully created');
      this.router.navigate(['/admin'])
        .catch(err => this.errorHandler.handle(err));
    } else {
      this.toastNotification.error(this.codesService.getMessageByDescription(
        response.data.description
      ));
    }
  }

  ngOnInit(): void {
  }
}
