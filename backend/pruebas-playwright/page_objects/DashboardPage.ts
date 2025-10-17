import { Page, expect } from '@playwright/test';

export class DashboardPage {
  readonly page: Page;
  readonly menuBeneficios;
  readonly opcionPremios;

  constructor(page: Page) {
    this.page = page;
    this.menuBeneficios = page.locator('shell-leal-text:has-text("Beneficios")').first();
    this.opcionPremios = page.locator('li:has-text("Premios")');
  }

  async hoverBeneficiosYClickPremios() {
    // 1️⃣ Localizar "Beneficios"
    const beneficios = this.page.locator('shell-leal-text >> text=Beneficios').first();
    await beneficios.waitFor({ state: 'visible', timeout: 5000 });
    await this.page.waitForTimeout(1000); // Pausa para que se vea

    // 2️⃣ Hover sobre "Beneficios"
    await beneficios.hover();
    await this.page.waitForTimeout(1000);

    // 3️⃣ Esperar que aparezca el submenú
    const dropdown = this.page.locator('.c-subitems-module');
    await dropdown.waitFor({ state: 'visible', timeout: 5000 });
    await this.page.waitForTimeout(1000);

    // 4️⃣ Localizar "Premios"
    const premios = dropdown.locator('li.c-subitem', { hasText: 'Premios' });
    await premios.waitFor({ state: 'visible', timeout: 5000 });
    await this.page.waitForTimeout(500);

    // 5️⃣ Click en "Premios"
    await premios.click();
    await this.page.waitForTimeout(1000); // Pausa final para que cargue
  }
}