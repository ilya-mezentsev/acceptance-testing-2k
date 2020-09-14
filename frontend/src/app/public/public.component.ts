import {Component, Inject, OnInit} from '@angular/core';
import {ErrorHandlerService} from '../services/errors/error-handler.service';
import {SessionStorageService} from '../services/session/session-storage.service';
import {Fetcher} from '../interfaces/fetcher';
import {Router} from '@angular/router';
import {ResponseStatus} from '../services/fetcher/statuses';

@Component({
  selector: 'app-public',
  templateUrl: './public.component.html',
  styleUrls: ['./public.component.scss']
})
export class PublicComponent implements OnInit {

  constructor(
    private readonly router: Router,
    private readonly errorHandler: ErrorHandlerService,
    private readonly sessionStorage: SessionStorageService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  ngOnInit(): void {
    this.fetcher.get('session/')
      .then(r => {
        if (r.status === ResponseStatus.OK) {
          this.sessionStorage.saveSession(r.data);
          return this.router.navigate(['/admin']);
        }
      })
      .catch(err => this.errorHandler.handle(err));
  }
}
