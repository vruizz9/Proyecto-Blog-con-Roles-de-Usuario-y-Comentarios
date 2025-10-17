import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BlogCards } from './blogCards';

describe('BlogCards', () => {
  let component: BlogCards;
  let fixture: ComponentFixture<BlogCards>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BlogCards]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BlogCards);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
