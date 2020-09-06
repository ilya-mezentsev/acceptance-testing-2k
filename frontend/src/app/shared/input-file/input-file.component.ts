import {Component, Input, OnInit, Output, EventEmitter} from '@angular/core';

@Component({
  selector: 'app-input-file',
  templateUrl: './input-file.component.html',
  styleUrls: ['./input-file.component.scss']
})
export class InputFileComponent implements OnInit {
  @Input() public label = '';
  @Input() public accept = '';
  @Output() public fileChanged = new EventEmitter<any>();

  constructor() { }

  ngOnInit(): void {
  }
}
