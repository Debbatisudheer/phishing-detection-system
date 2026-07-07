import { test, expect } from "@playwright/test";

test(
  "Session Expiry Test",
  async ({ page }) => {

    await page.goto("/login");

    await page
      .getByPlaceholder("Username")
      .fill("sudheer");

    await page
      .getByPlaceholder("Password")
      .fill("password123");

    page.on("dialog", async (dialog) => {
      await dialog.accept();
    });

    await page
      .getByRole("button", {
        name: "Login",
      })
      .click();

    await expect(page).toHaveURL("/");

    await expect(
      page.getByRole("heading", {
        name: "SOC Dashboard",
      })
    ).toBeVisible();

    // Simulate session expiry
    await page.evaluate(() => {
      localStorage.removeItem("token");
    });

    // Refresh the page
    await page.reload();

    // ProtectedRoute should redirect to Login
    await expect(page).toHaveURL(/login/);

    await expect(
      page.getByRole("heading", {
        name: "Login",
      })
    ).toBeVisible();

  },
);