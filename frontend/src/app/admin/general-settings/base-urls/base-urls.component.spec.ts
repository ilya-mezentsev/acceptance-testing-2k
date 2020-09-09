import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BaseUrlsComponent } from './base-urls.component';

describe('BaseUrlsComponent', () => {
  let component: BaseUrlsComponent;
  let fixture: ComponentFixture<BaseUrlsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BaseUrlsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BaseUrlsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
