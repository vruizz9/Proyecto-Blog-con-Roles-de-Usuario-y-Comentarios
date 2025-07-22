import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TarjetaBlog } from './tarjeta-blog';

describe('TarjetaBlog', () => {
  let component: TarjetaBlog;
  let fixture: ComponentFixture<TarjetaBlog>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TarjetaBlog]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TarjetaBlog);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
