import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CommandFieldsComponent } from './command-fields.component';

describe('CommandFieldsComponent', () => {
  let component: CommandFieldsComponent;
  let fixture: ComponentFixture<CommandFieldsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CommandFieldsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CommandFieldsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
