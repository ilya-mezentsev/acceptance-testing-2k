import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { RunTestsComponent } from './run-tests.component';

describe('RunTestsComponent', () => {
  let component: RunTestsComponent;
  let fixture: ComponentFixture<RunTestsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ RunTestsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(RunTestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
