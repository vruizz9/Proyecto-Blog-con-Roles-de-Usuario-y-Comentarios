import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CajaComentario } from './caja-comentario';

describe('CajaComentario', () => {
  let component: CajaComentario;
  let fixture: ComponentFixture<CajaComentario>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CajaComentario]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CajaComentario);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
