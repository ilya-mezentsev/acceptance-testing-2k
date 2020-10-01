import { Component, OnInit, Input } from '@angular/core';

export type Link = {path: string, name: string};

@Component({
  selector: 'app-sidenav-links',
  templateUrl: './sidenav-links.component.html',
  styleUrls: ['./sidenav-links.component.scss']
})
export class SidenavLinksComponent implements OnInit {
  @Input() public links: Link[] = [];

  constructor() { }

  ngOnInit(): void {
  }
}
