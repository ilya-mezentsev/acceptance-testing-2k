import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import {KeyValueMapping} from '../../types/types';

@Component({
  selector: 'app-cookies-adding',
  templateUrl: './cookies-adding.component.html',
  styleUrls: ['./cookies-adding.component.scss']
})
export class CookiesAddingComponent implements OnInit {
  public noActiveBlocks = -1;
  public activeCookieIndex = this.noActiveBlocks;

  @Input() public cookies: KeyValueMapping[] = [];
  @Output() public addCookie = new EventEmitter();
  @Output() public removeCookie = new EventEmitter<string>();

  constructor() { }

  public hasCookies(): boolean {
    return this.cookies.length > 0;
  }

  public shouldBlurCookie(index: number): boolean {
    return this.activeCookieIndex !== this.noActiveBlocks && index !== this.activeCookieIndex;
  }

  ngOnInit(): void {
  }
}
