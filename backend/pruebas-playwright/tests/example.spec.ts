import { test } from '@playwright/test';
import { LoginPage } from '../page_objects/LoginPage';
import { DashboardPage } from '../page_objects/DashboardPage';
import { PremiosPage } from '../page_objects/PremiosPage';

test('Flujo completo - Login y Creación de premios', async ({ page }) => {
  const loginPage = new LoginPage(page);
  const dashboardPage = new DashboardPage(page);
  const premiosPage = new PremiosPage(page);

  // 1️⃣ Ir al login
  await loginPage.goto();

  // 2️⃣ Iniciar sesión
  await loginPage.login('admin.12@puntosleal.com', 'Leal2024*');

  // 3️⃣ Esperar dashboard con pausa visual
  await page.waitForURL('**/dashboard/inicio', { timeout: 10000 });
  await page.waitForTimeout(1500); // Igual que en DashboardPage

  // 4️⃣ Hover y click en "Premios"
  await dashboardPage.hoverBeneficiosYClickPremios();

  // 5️⃣ Esperar que cargue premios
  await page.waitForURL('**/dashboard/beneficios/premios', { timeout: 10000 });
  await page.waitForTimeout(1500); // Pequeña pausa final para estabilidad

  await premiosPage.clickCrearPremios();

});

