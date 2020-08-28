import {Component, Input, OnInit, Output, EventEmitter} from '@angular/core';

@Component({
  selector: 'app-button-with-icon',
  templateUrl: './button-with-icon.component.html',
  styleUrls: ['./button-with-icon.component.scss']
})
export class ButtonWithIconComponent implements OnInit {
  @Input() public iconName = '';
  @Input() public iconPosition = '';
  @Input() public text = '';
  @Input() public small = false;
  @Output() public clicked = new EventEmitter<boolean>();

  constructor() { }

  ngOnInit(): void {
  }
}
