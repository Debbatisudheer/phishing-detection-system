import { test } from "@playwright/test";
import { LoginPage } from "../pages/LoginPage";

test(
    "Login Page Smoke Test",
    async ({ page }) => {

        const loginPage =
            new LoginPage(page);

        await loginPage.goto();

        await loginPage.verifyLoginPage();

       await loginPage.login(
    "sudheer",
    "password123",
);

    },
);