import { Routes } from '@angular/router';
import { LoginComponent } from './ui/auth/login/login';
import { List } from './ui/blog/list/list';
import { DetailComponent } from './ui/blog/detail/detail';
import { Create } from './ui/blog/create/create';
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
      { path: 'inicio', component: List },
      { path: 'blog/:id', component: DetailComponent },
      { path : 'crear-blog', component: Create}
    ]
  },
  {
    path: '**',
    redirectTo: 'login'
  }
];
