// Content managed by Project Forge, see [projectforge.md] for details.
import { test, expect } from '@playwright/test';

test.describe("pages", () => {
  test('home', async ({ page }) => {
    await page.goto("/");
    await expect(page).toHaveTitle("Test Project");
  });

  test('about', async ({ page }) => {
    await page.goto("/about");
  });
});
