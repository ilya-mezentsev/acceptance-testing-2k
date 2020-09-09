import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-input-checkbox',
  templateUrl: './input-checkbox.component.html',
  styleUrls: ['./input-checkbox.component.scss']
})
export class InputCheckboxComponent implements OnInit {
  @Input() public value = false;
  @Input() public label = '';
  @Output() public valueEmitter = new EventEmitter<boolean>();

  public changed(): void {
    this.value = !this.value;
    this.valueEmitter.emit(this.value);
  }

  constructor() { }

  ngOnInit(): void {
  }
}
