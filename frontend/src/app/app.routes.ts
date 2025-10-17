import { Routes } from '@angular/router';
import { LoginComponent } from './ui/auth/login/login';
import { List } from './ui/blog/list/list';
import { DetailComponent } from './ui/blog/detail/detail';
import { LayoutComponent } from './shared/componentes/layout/layout';

export const routes: Routes = [
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full'
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: '',
    component: LayoutComponent,
    children: [
      { path: 'lista', component: List },
      { path: 'blog/:id', component: DetailComponent }
    ]
  },
  {
    path: '**',
    redirectTo: 'login'
  }
];
