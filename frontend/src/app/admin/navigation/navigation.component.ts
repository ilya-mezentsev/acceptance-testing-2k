import {Component, Inject, OnInit} from '@angular/core';
import {Fetcher} from '../../interfaces/fetcher';
import {ErrorHandlerService} from '../../services/errors/error-handler.service';
import {Router} from '@angular/router';
import {SessionStorageService} from '../../services/session/session-storage.service';
import {MaterializeInitService} from '../../services/materialize/materialize-init.service';
import {StorageService} from '../services/storage/storage.service';
import {Link} from '../../shared/sidenav-links/sidenav-links.component';

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent implements OnInit {
  public readonly mainRoutes: Link[] = [
    {path: './objects-list', name: 'Objects list'},
    {path: './create-object', name: 'Create object'},
    {path: './run-tests', name: 'Run tests'},
  ];
  public readonly generalSettingsRoutes: Link[] = [
    {path: './general-base-urls', name: 'Base URLs'},
    {path: './general-timeouts', name: 'Timeouts'},
    {path: './general-headers', name: 'Headers'},
    {path: './general-cookies', name: 'Cookies'},
  ];
  public readonly infoRoutes: Link[] = [
    {path: './faq', name: 'FAQ'},
    {path: './about', name: 'About'},
  ];

  constructor(
    private readonly router: Router,
    private readonly storage: StorageService,
    private readonly materializeInit: MaterializeInitService,
    private readonly sessionStorage: SessionStorageService,
    private readonly errorHandler: ErrorHandlerService,
    @Inject('Fetcher') private readonly fetcher: Fetcher,
  ) { }

  public signOut(): void {
    this.fetcher.delete('/session')
      .then(() => this.sessionStorage.deleteSession())
      .then(() => this.storage.invalidate())
      .then(() => this.router.navigate(['']))
      .catch(err => this.errorHandler.handle(err));
  }

  ngOnInit(): void {
    this.materializeInit.initSidenav();
    this.materializeInit.initDropdowns();
  }
}
