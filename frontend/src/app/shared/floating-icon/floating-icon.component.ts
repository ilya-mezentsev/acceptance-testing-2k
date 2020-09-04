import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-floating-icon',
  templateUrl: './floating-icon.component.html',
  styleUrls: ['./floating-icon.component.scss']
})
export class FloatingIconComponent implements OnInit {
  @Input() public iconName = '';

  constructor() { }

  ngOnInit(): void {
  }
}
