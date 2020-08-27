import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-button-modal-trigger',
  templateUrl: './button-modal-trigger.component.html',
  styleUrls: ['./button-modal-trigger.component.scss']
})
export class ButtonModalTriggerComponent implements OnInit {
  @Input() public text = '';
  @Input() public dataTarget = '';

  constructor() { }

  ngOnInit(): void {
  }
}
