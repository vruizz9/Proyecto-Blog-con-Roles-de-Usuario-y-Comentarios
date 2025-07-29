import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './login.html',
  styleUrls: ['./login.scss']
})
export class LoginComponent {
  username: string = '';
  password: string = '';
  mostrarContrasena: boolean = false;

  toggleMostrarContrasena() {
    this.mostrarContrasena = !this.mostrarContrasena;
  }

  onLogin(form: any) {
    if (form.valid) {
      console.log('Usuario:', this.username);
      console.log('Contraseña:', this.password);
      // Aquí puedes implementar tu lógica de autenticación
    } else {
      console.warn('Formulario inválido');
    }
  }
}
