import {Component, Input, OnInit, Output, EventEmitter} from '@angular/core';

@Component({
  selector: 'app-button',
  templateUrl: './button.component.html',
  styleUrls: ['./button.component.scss']
})
export class ButtonComponent implements OnInit {
  @Input() public isDisabled = false;
  @Input() public text = '';
  @Output() public clicked = new EventEmitter<null>();

  ngOnInit(): void {
  }
}
