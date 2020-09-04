import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateCommandComponent } from './create-command.component';

describe('CreateCommandComponent', () => {
  let component: CreateCommandComponent;
  let fixture: ComponentFixture<CreateCommandComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CreateCommandComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateCommandComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
