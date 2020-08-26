import {Component, Inject, OnInit} from '@angular/core';
import {Fetcher} from '../../interfaces/fetcher';
import {ErrorHandlerService} from '../../services/errors/error-handler.service';
import {Router} from '@angular/router';
import {SessionStorageService} from '../../services/session/session-storage.service';

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent implements OnInit {

  constructor(
    private readonly router: Router,
    private readonly sessionStorage: SessionStorageService,
    private readonly errorHandler: ErrorHandlerService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  public signOut(): void {
    this.fetcher.delete('/session/')
      .then(() => this.sessionStorage.deleteSession())
      .then(() => this.router.navigate(['/']))
      .catch(err => this.errorHandler.handle(err));
  }

  public get login(): string {
    return this.sessionStorage.getSessionLogin();
  }

  ngOnInit(): void {
  }
}
