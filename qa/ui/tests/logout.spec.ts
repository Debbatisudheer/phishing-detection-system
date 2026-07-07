import { test, expect } from "@playwright/test";

test(
  "Logout Test",
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

    await page
      .getByRole("button", {
        name: "Logout",
      })
      .click();

    await expect(page).toHaveURL(/login/);

    await expect(
      page.getByRole("heading", {
        name: "Login",
      })
    ).toBeVisible();

    const token =
      await page.evaluate(() =>
        localStorage.getItem("token")
      );

    expect(token).toBeNull();

  },
);