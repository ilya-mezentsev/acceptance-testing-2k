import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ButtonModalTriggerComponent } from './button-modal-trigger.component';

describe('ButtonModalTriggerComponent', () => {
  let component: ButtonModalTriggerComponent;
  let fixture: ComponentFixture<ButtonModalTriggerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ButtonModalTriggerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ButtonModalTriggerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
