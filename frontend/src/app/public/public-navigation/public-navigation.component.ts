import { Component, OnInit } from '@angular/core';
import {MaterializeInitService} from '../../services/materialize/materialize-init.service';

@Component({
  selector: 'app-public-navigation',
  templateUrl: './public-navigation.component.html',
  styleUrls: ['./public-navigation.component.scss']
})
export class PublicNavigationComponent implements OnInit {

  constructor(
    private readonly materializeInit: MaterializeInitService,
  ) { }

  ngOnInit(): void {
    this.materializeInit.initSidenav();
  }
}
