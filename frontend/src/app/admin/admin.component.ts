import {Component, Inject, OnInit} from '@angular/core';
import {ErrorResponse, Fetcher, Response} from '../interfaces/fetcher';
import {ResponseStatus} from '../services/fetcher/statuses';
import {TestCommand, TestObject} from './types/types';
import {NavigationEnd, Router} from '@angular/router';
import {ErrorHandlerService} from '../services/errors/error-handler.service';
import {ToastNotificationService} from '../services/notification/toast-notification.service';
import {CodesService} from './services/errors/codes.service';
import {SessionStorageService} from '../services/session/session-storage.service';
import {StorageService} from './services/storage/storage.service';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent implements OnInit {
  constructor(
    private readonly router: Router,
    private readonly errorHandler: ErrorHandlerService,
    private readonly toastNotification: ToastNotificationService,
    private readonly codesService: CodesService,
    private readonly sessionStorage: SessionStorageService,
    private readonly storage: StorageService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  ngOnInit(): void {
    this.fetchData();

    // we need to fetch data again because navigation is performed after some data changes
    this.router.events.subscribe(event => event instanceof NavigationEnd && this.fetchData());
  }

  private fetchData(): void {
    this.fetchTestObjects();
    this.fetchTestCommands();
  }

  private fetchTestObjects(): void {
    if (this.storage.hasObjects()) {
      return;
    }

    this.fetcher.get(`entity/test-object/${this.sessionStorage.getSessionId()}/`)
      .then(r => {
        if (r.status === ResponseStatus.OK) {
          this.storage.objects = (r as Response<TestObject[]>).data || [];
        } else {
          this.sendErrorNotification(r as ErrorResponse);
        }
      })
      .catch(err => this.errorHandler.handle(err));
  }

  private fetchTestCommands(): void {
    if (this.storage.hasCommands()) {
      return;
    }

    this.fetcher.get(`entity/test-command/${this.sessionStorage.getSessionId()}/`)
      .then(r => {
        if (r.status === ResponseStatus.OK) {
          this.storage.commands = (r as Response<TestCommand[]>).data || [];
        } else {
          this.sendErrorNotification(r as ErrorResponse);
        }
      })
      .catch(err => this.errorHandler.handle(err));
  }

  private sendErrorNotification(error: ErrorResponse): void {
    this.toastNotification.error(this.codesService.getMessageByDescription(
      error.data.description
    ));
  }
}
