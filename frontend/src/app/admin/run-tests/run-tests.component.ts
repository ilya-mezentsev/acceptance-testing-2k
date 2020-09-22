import {Component, Inject, OnInit, EventEmitter} from '@angular/core';
import {ErrorResponse, FileSender, Response} from '../../interfaces/fetcher';
import {ErrorHandlerService} from '../../services/errors/error-handler.service';
import {ToastNotificationService} from '../../services/notification/toast-notification.service';
import {CreateTestsFile, RunTestsResult, TestsReport, ServiceError} from '../types/types';
import {ResponseStatus} from '../../services/fetcher/statuses';
import {CodesService} from '../services/errors/codes.service';
import {MaterializeInitService} from '../../services/materialize/materialize-init.service';
import {TestsRunnerService} from '../services/tests-runner/tests-runner.service';

@Component({
  selector: 'app-run-tests',
  templateUrl: './run-tests.component.html',
  styleUrls: ['./run-tests.component.scss']
})
export class RunTestsComponent implements OnInit {
  private file: File;
  public testsReport: TestsReport;
  public awaitingTestsResults = false;
  public hasTestsReport = false;
  public readonly resetInput = new EventEmitter();
  private readonly maxFileSize = 32 * 1024 * 1024;  // 32 MB

  constructor(
    private readonly codesService: CodesService,
    private readonly testsRunner: TestsRunnerService,
    private readonly errorHandler: ErrorHandlerService,
    private readonly toastNotification: ToastNotificationService,
    private readonly materializeInit: MaterializeInitService,
    @Inject('FileSender') private readonly fileSender: FileSender,
  ) { }

  public fileChange(event): void {
    if (event?.target?.files?.length > 0) {
      this.file = event.target.files[0];
    } else {
      this.file = null;
    }
  }

  public get hasFile(): boolean {
    return !!this.file;
  }

  public get hasErrors(): boolean {
    return (this.testsReport?.errors || []).length > 0;
  }

  public getTestCaseRows(testCaseText: string): string[] {
    return testCaseText.split('\n').map(s => s.trim());
  }

  public runTests(): void {
    if (!this.hasFile) {
      this.toastNotification.info('You need to choose file first');
      return;
    } else if (this.file.size > this.maxFileSize) {
      this.toastNotification.error('File is too large');
      return;
    }

    const fd = new FormData();
    fd.append('tests_cases_file', this.file);

    this.awaitingTestsResults = true;
    this.hasTestsReport = false;

    this.fileSender
      .sendFile<CreateTestsFile>('tests-file', fd)
      .then(r => {
        this.processCreateTestsFileResponse(r);
        this.resetInput.emit();
      })
      .catch(err => {
        this.awaitingTestsResults = false;
        this.errorHandler.handle(err);
      });
  }

  private processCreateTestsFileResponse(
    response: Response<CreateTestsFile> | ErrorResponse
  ): void {
    if (response.status === ResponseStatus.OK) {
      this.openConnectionForRunTests(
        (response.data as CreateTestsFile).filename
      );
    } else {
      this.awaitingTestsResults = false;
      this.toastNotification.error(this.codesService.getMessageByDescription(
        (response as ErrorResponse).data.description
      ));
    }
  }

  private openConnectionForRunTests(filename: string): void {
    this.testsRunner.run<RunTestsResult | ServiceError>(filename)
      .then(r => this.processTestsReport(r))
      .catch(err => this.errorHandler.handle(err))
      .finally(() => this.awaitingTestsResults = false);
  }

  private processTestsReport(response: RunTestsResult | ServiceError): void {
    if (!!(response as RunTestsResult).report) {
      this.testsReport = (response as RunTestsResult).report;
      this.hasTestsReport = true;

      setTimeout(() => this.materializeInit.initTooltips(), 0);
    } else if (!!(response as RunTestsResult).applicationError) {
      this.toastNotification.error(
        this.codesService.getMessageByDescription(
          (response as RunTestsResult).applicationError.description
        )
      );
    } else {
      this.toastNotification.error(
        this.codesService.getMessageByDescription((response as ServiceError).description)
      );
    }
  }

  ngOnInit(): void {
  }
}
