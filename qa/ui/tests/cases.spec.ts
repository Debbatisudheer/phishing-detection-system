import { test } from "@playwright/test";
import { CasesPage } from "../pages/CasesPage";

test(
    "Cases Page Smoke Test",
    async ({ page }) => {

        page.on(
            "dialog",
            async dialog => {
                await dialog.accept();
            },
        );

        await page.goto("/login");

        await page
            .getByPlaceholder("Username")
            .fill("sudheer");

        await page
            .getByPlaceholder("Password")
            .fill("password123");

        await Promise.all([
            page.waitForURL("**/"),
            page.getByRole("button", {
                name: "Login",
            }).click(),
        ]);

        const cases =
            new CasesPage(page);

        await cases.goto();

        await cases.verifyCasesLoaded();

        await cases.verifyAnalystVisible();

        await cases.verifyStatusVisible();

        await cases.verifyNotesVisible();

    },
);