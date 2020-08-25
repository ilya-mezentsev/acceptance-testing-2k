import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-router-button',
  templateUrl: './router-button.component.html',
  styleUrls: ['./router-button.component.scss']
})
export class RouterButtonComponent implements OnInit {
  @Input() public text = '';

  ngOnInit(): void {
  }
}
