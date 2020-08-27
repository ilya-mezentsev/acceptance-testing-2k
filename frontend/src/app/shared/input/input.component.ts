import {Component, Input, OnInit, Output, EventEmitter} from '@angular/core';
import {Md5} from 'ts-md5';

@Component({
  selector: 'app-input',
  templateUrl: './input.component.html',
  styleUrls: ['./input.component.scss']
})
export class InputComponent implements OnInit {
  public id = '';
  @Input() public value: any;
  @Input() public readonly label: string;
  @Input() public readonly type: string;
  @Output() public valueEmitter = new EventEmitter<string>();

  constructor() { }

  public valueChanged(): void {
    this.valueEmitter.emit(this.value);
  }

  ngOnInit(): void {
    this.id = `${Md5.hashStr((new Date()).getMilliseconds().toString())}${Math.random()}`;
  }
}
