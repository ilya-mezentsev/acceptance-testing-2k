import { Component, OnInit } from '@angular/core';
import {MaterializeInitService} from '../../services/materialize/materialize-init.service';

@Component({
  selector: 'app-faq',
  templateUrl: './faq.component.html',
  styleUrls: ['./faq.component.scss']
})
export class FaqComponent implements OnInit {

  constructor(
    private readonly materializeInit: MaterializeInitService,
  ) { }

  ngOnInit(): void {
    this.materializeInit.initCollapsibleWithoutOptions();
  }
}
