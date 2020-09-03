import {Component, OnInit} from '@angular/core';
import {TestObject} from '../types/types';
import {StorageService} from '../services/storage/storage.service';

@Component({
  selector: 'app-objects-list',
  templateUrl: './objects-list.component.html',
  styleUrls: ['./objects-list.component.scss']
})
export class ObjectsListComponent implements OnInit {
  constructor(
    private readonly storage: StorageService,
  ) { }

  public hasObjects(): boolean {
    return this.storage.objects.length > 0;
  }

  public getObjects(): TestObject[] {
    return this.storage.objects;
  }

  ngOnInit(): void {
  }
}
