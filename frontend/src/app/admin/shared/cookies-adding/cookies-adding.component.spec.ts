import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CookiesAddingComponent } from './cookies-adding.component';

describe('CookiesAddingComponent', () => {
  let component: CookiesAddingComponent;
  let fixture: ComponentFixture<CookiesAddingComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CookiesAddingComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CookiesAddingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
