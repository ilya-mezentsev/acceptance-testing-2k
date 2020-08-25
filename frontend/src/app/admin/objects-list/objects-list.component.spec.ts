import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ObjectsListComponent } from './objects-list.component';

describe('ObjectsListComponent', () => {
  let component: ObjectsListComponent;
  let fixture: ComponentFixture<ObjectsListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ObjectsListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ObjectsListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
