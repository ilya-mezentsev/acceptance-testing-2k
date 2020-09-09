import {Component, Input, OnInit, Output, EventEmitter} from '@angular/core';
import {KeyValueMapping} from '../../types/types';

@Component({
  selector: 'app-headers-adding',
  templateUrl: './headers-adding.component.html',
  styleUrls: ['./headers-adding.component.scss']
})
export class HeadersAddingComponent implements OnInit {
  public noActiveBlocks = -1;
  public activeHeaderIndex = this.noActiveBlocks;

  @Input() public headers: KeyValueMapping[] = [];
  @Output() public addHeader = new EventEmitter();
  @Output() public removeHeader = new EventEmitter<string>();

  constructor() { }

  public hasHeaders(): boolean {
    return this.headers.length > 0;
  }

  public shouldBlurHeader(index: number): boolean {
    return this.activeHeaderIndex !== this.noActiveBlocks && index !== this.activeHeaderIndex;
  }

  ngOnInit(): void {
  }
}
