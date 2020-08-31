import {Component, OnInit, Output, EventEmitter, Input} from '@angular/core';
import {KeyValueMapping, TestCommandSettings} from "../../../types/types";

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
  @Output() public removeHeader = new EventEmitter<number>();
  @Output() public addCookie = new EventEmitter();
  @Output() public removeCookie = new EventEmitter<number>();

  public hasHeaders(): boolean {
    return this.headers.length > 0;
  }

  public hasCookies(): boolean {
    return this.cookies.length > 0;
  }

  constructor() { }

  ngOnInit(): void {
  }
}
