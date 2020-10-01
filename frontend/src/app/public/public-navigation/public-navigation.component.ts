import { Component, OnInit } from '@angular/core';
import {MaterializeInitService} from '../../services/materialize/materialize-init.service';
import {Link} from '../../shared/sidenav-links/sidenav-links.component';

@Component({
  selector: 'app-public-navigation',
  templateUrl: './public-navigation.component.html',
  styleUrls: ['./public-navigation.component.scss']
})
export class PublicNavigationComponent implements OnInit {
  public readonly routes: Link[] = [
    {path: './sign-in', name: 'Sign in'},
    {path: './sign-up', name: 'Sign up'},
    {path: './about', name: 'About'},
  ];

  constructor(
    private readonly materializeInit: MaterializeInitService,
  ) { }

  ngOnInit(): void {
    this.materializeInit.initSidenav();
  }
}
