import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HeadersAddingComponent } from './headers-adding.component';

describe('HeadersAddingComponent', () => {
  let component: HeadersAddingComponent;
  let fixture: ComponentFixture<HeadersAddingComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HeadersAddingComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HeadersAddingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
