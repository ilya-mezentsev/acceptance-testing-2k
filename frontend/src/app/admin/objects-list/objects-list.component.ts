import {Component, Inject, OnInit} from '@angular/core';
import {ErrorHandlerService} from "../../services/errors/error-handler.service";
import {ToastNotificationService} from "../../services/notification/toast-notification.service";
import {CodesService} from "../services/errors/codes.service";
import {ErrorResponse, Fetcher, Response} from "../../interfaces/fetcher";
import {Router} from "@angular/router";
import {SessionStorageService} from "../../services/session/session-storage.service";
import {Command, Object} from "../types/types";
import {ResponseStatus} from "../../services/fetcher/statuses";
import {StorageService} from "../services/storage/storage.service";

@Component({
  selector: 'app-objects-list',
  templateUrl: './objects-list.component.html',
  styleUrls: ['./objects-list.component.scss']
})
export class ObjectsListComponent implements OnInit {
  constructor(
    private readonly router: Router,
    private readonly errorHandler: ErrorHandlerService,
    private readonly toastNotification: ToastNotificationService,
    private readonly codesService: CodesService,
    private readonly sessionStorage: SessionStorageService,
    private readonly storage: StorageService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  public hasObjects(): boolean {
    return this.storage.hasObjects();
  }

  public getObjects(): Object[] {
    return this.storage.objects;
  }

  public getObjectCommands(objectName: string): Command[] {
    return this.storage.commands.filter(c => c.object_name == objectName);
  }

  ngOnInit(): void {
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
          this.storage.objects = (r as Response<Object[]>).data || [];
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
          this.storage.commands = (r as Response<Command[]>).data || [];
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
