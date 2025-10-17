import { expect, Page } from '@playwright/test';

export class PremiosPage {
  private page: Page;
  private crearPremiosButton;

  constructor(page: Page) {
    this.page = page;
    this.crearPremiosButton = page.locator('leal-button-ds[type="primary"]')
  .filter({ hasText: 'Crear premios' });
  }

  async clickCrearPremios() {
    await this.crearPremiosButton.click();
  }
}
