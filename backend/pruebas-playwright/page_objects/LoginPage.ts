import { expect, Page } from '@playwright/test';

export class LoginPage {
  private page: Page;
  private emailInput;
  private passwordInput;
  private loginButton;

  constructor(page: Page) {
    this.page = page;
    this.emailInput = page.getByRole('textbox', { name: 'Escribe tu correo electrónico' });
    this.passwordInput = page.getByRole('textbox', { name: 'Escribe tu contraseña' });
    this.loginButton = page.locator('button.btn.btn-primary.rounded-both.w-full');
  }

  async goto() {
    await this.page.goto('https://qalogin.leal.co/');
    await this.page.waitForTimeout(1000); // Pausa para cargar bien la página
  }

  async login(email: string, password: string) {
    // Esperar que el campo email esté visible
    await this.emailInput.waitFor({ state: 'visible', timeout: 5000 });
    await this.emailInput.fill(email);
    await this.page.waitForTimeout(800); // Pausa visual

    // Esperar que el campo password esté visible
    await this.passwordInput.waitFor({ state: 'visible', timeout: 5000 });
    await this.passwordInput.fill(password);
    await this.page.waitForTimeout(800); // Pausa visual

    // Esperar que el botón esté listo
    await this.loginButton.waitFor({ state: 'visible', timeout: 10000 });
    await expect(this.loginButton).toBeEnabled();
    await this.page.waitForTimeout(500); // Pequeña pausa

    // Clic en el botón con fallback
    try {
      await this.loginButton.click({ force: true });
    } catch {
      await this.page.evaluate(() => {
        document.querySelector<HTMLButtonElement>(
          'button.btn.btn-primary.rounded-both.w-full'
        )?.click();
      });
    }

    await this.page.waitForTimeout(1500); // Espera para transición
    await this.page.waitForURL('**/dashboard/**', { timeout: 10000 });
  }
}