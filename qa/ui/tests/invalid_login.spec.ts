import { test, expect } from "@playwright/test";

test(
  "Invalid Login Test",
  async ({ page }) => {

    await page.goto("/login");

    await page
      .getByPlaceholder("Username")
      .fill("invaliduser");

    await page
      .getByPlaceholder("Password")
      .fill("wrongpassword");

    page.on("dialog", async (dialog) => {

      expect(dialog.message()).toBe(
        "Invalid Username or Password"
      );

      await dialog.accept();

    });

    await page
      .getByRole("button", {
        name: "Login",
      })
      .click();

    await expect(page).toHaveURL(
      /login/
    );

  },
);