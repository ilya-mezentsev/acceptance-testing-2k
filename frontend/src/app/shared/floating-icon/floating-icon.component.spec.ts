import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FloatingIconComponent } from './floating-icon.component';

describe('FloatingIconComponent', () => {
  let component: FloatingIconComponent;
  let fixture: ComponentFixture<FloatingIconComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FloatingIconComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FloatingIconComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
