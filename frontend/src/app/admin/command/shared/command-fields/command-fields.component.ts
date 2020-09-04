import {Component, OnInit, Output, EventEmitter, Input} from '@angular/core';
import {KeyValueMapping, TestCommandSettings} from '../../../types/types';
import {MaterializeInitService} from '../../../../services/materialize/materialize-init.service';

@Component({
  selector: 'app-command-fields',
  templateUrl: './command-fields.component.html',
  styleUrls: ['./command-fields.component.scss']
})
export class CommandFieldsComponent implements OnInit {
  public noActiveBlocks = -1;
  public activeHeaderIndex = this.noActiveBlocks;
  public activeCookieIndex = this.noActiveBlocks;

  public passArgumentsInURL = false;
  @Input() public headers: KeyValueMapping[] = [];
  @Input() public cookies: KeyValueMapping[] = [];
  @Input() public commandSettings: TestCommandSettings;
  @Output() public addHeader = new EventEmitter();
  @Output() public removeHeader = new EventEmitter<string>();
  @Output() public addCookie = new EventEmitter();
  @Output() public removeCookie = new EventEmitter<string>();

  public hasHeaders(): boolean {
    return this.headers.length > 0;
  }

  public hasCookies(): boolean {
    return this.cookies.length > 0;
  }

  public shouldBlurHeader(index: number): boolean {
    return this.activeHeaderIndex !== this.noActiveBlocks && index !== this.activeHeaderIndex;
  }

  public shouldBlurCookie(index: number): boolean {
    return this.activeCookieIndex !== this.noActiveBlocks && index !== this.activeCookieIndex;
  }

  constructor(
    private readonly materializeInit: MaterializeInitService,
  ) { }

  ngOnInit(): void {
    this.materializeInit.initSelects();
  }
}
