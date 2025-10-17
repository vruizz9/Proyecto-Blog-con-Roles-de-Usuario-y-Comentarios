import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { NavbarComponent } from '../navbar/navbar'; // ajusta la ruta según tu estructura

@Component({
  selector: 'app-layout',
  standalone: true,
  imports: [RouterOutlet, NavbarComponent],
  templateUrl: './layout.html',
  styleUrls: ['./layout.scss']
})

export class LayoutComponent {}