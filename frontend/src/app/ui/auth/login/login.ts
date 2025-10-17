import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, NgForm } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../../core/application/auth/auth.service';

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
  errorMessage: string = '';

  constructor(private router: Router, private authService: AuthService) {}

  toggleMostrarContrasena(): void {
    this.mostrarContrasena = !this.mostrarContrasena;
  }

  onLogin(form: NgForm): void {
    if (form.valid) {
      this.authService.login(this.username, this.password).subscribe(user => {
        if (user) {
          console.log('✅ Login exitoso:', user);
          localStorage.setItem('currentUser', JSON.stringify(user));
          this.router.navigate(['/lista']);
        } else {
          this.errorMessage = 'Usuario o contraseña incorrectos';
        }
      });
    } else {
      this.errorMessage = 'Formulario inválido';
      form.control.markAllAsTouched();
    }
  }
}