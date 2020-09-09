import {Component, OnInit, Output, EventEmitter, Input} from '@angular/core';
import {KeyValueMapping, TestCommandSettings} from '../../../types/types';
import {MaterializeInitService} from '../../../../services/materialize/materialize-init.service';

@Component({
  selector: 'app-command-fields',
  templateUrl: './command-fields.component.html',
  styleUrls: ['./command-fields.component.scss']
})
export class CommandFieldsComponent implements OnInit {
  public passArgumentsInURL = false;
  @Input() public headers: KeyValueMapping[] = [];
  @Input() public cookies: KeyValueMapping[] = [];
  @Input() public commandSettings: TestCommandSettings;
  @Output() public addHeader = new EventEmitter();
  @Output() public removeHeader = new EventEmitter<string>();
  @Output() public addCookie = new EventEmitter();
  @Output() public removeCookie = new EventEmitter<string>();

  constructor(
    private readonly materializeInit: MaterializeInitService,
  ) { }

  ngOnInit(): void {
    this.materializeInit.initSelects();
  }
}
